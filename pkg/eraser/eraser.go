package eraser

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"git.linuxrocker.com/mattburchett/Housekeeper/pkg/config"
	"git.linuxrocker.com/mattburchett/Housekeeper/pkg/model"
)

// LookupMovieFileLocation will gather a list of Information based on IDs returned by locator.GetTitles
func LookupMovieFileLocation(config config.Config, ids []int) []string {
	fileList := make([]string, 0)

	for _, i := range ids {
		plexURL := fmt.Sprintf("%s:%d%s%d%s%s", config.PlexHost, config.PlexPort, "/library/metadata/", i, "/?X-Plex-Token=", config.PlexToken)

		req, err := http.NewRequest(http.MethodGet, plexURL, nil)

		httpClient := http.Client{}
		req.Header.Set("User-Agent", "Housekeeper")

		res, getErr := httpClient.Do(req)
		if getErr != nil {
			log.Fatal(getErr)
		}

		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}

		if err != nil {
			log.Fatal(err)
		}
		plexModel := model.XMLPlexMovieAPI{}
		xml.Unmarshal(body, &plexModel)
		fileList = append(fileList, filepath.Dir(plexModel.Video.Media.Part.File))
	}
	return fileList
}

// LookupTVFileLocation will gather a list of Information based on IDs returned by locator.GetTitles
func LookupTVFileLocation(config config.Config, ids []int) []string {
	fileList := make([]string, 0)
	m := make(map[string]bool)

	for _, i := range ids {
		plexURL := fmt.Sprintf("%s:%d%s%d%s%s", config.PlexHost, config.PlexPort, "/library/metadata/", i, "/allLeaves/?X-Plex-Token=", config.PlexToken)

		req, err := http.NewRequest(http.MethodGet, plexURL, nil)

		httpClient := http.Client{}
		req.Header.Set("User-Agent", "Housekeeper")

		res, getErr := httpClient.Do(req)
		if getErr != nil {
			log.Fatal(getErr)
		}

		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}

		if err != nil {
			log.Fatal(err)
		}
		plexModel := model.XMLPlexTVAPI{}
		xml.Unmarshal(body, &plexModel)

		plexTV := plexModel.Video
		results := make([]string, 0)

		for _, i := range plexTV {
			fileList = append(fileList, filepath.Dir(filepath.Dir(i.Media.Part.File)))
		}

		for _, r := range fileList {
			if _, ok := m[r]; !ok {
				m[r] = true
				results = append(results, r)
			}
		}

	}
	return results
}

func isValueInList(value string, list []string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

// DeleteMovies will actually perform the deletion.
func DeleteMovies(delete bool, files []string) error {
	var err error
	if delete {
		for _, i := range files {
			fmt.Printf("Removing %v\n", i)
			err = os.RemoveAll(i)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return err
}
