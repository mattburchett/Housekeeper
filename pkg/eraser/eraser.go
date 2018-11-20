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

// LookupFileLocation will gather a list of Information based on IDs returned by locator.GetTitles
func LookupFileLocation(config config.Config, ids []int) []string {
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
		plexModel := model.XMLPlexAPI{}
		xml.Unmarshal(body, &plexModel)
		fileList = append(fileList, filepath.Dir(plexModel.Video.Media.Part.File))
	}
	return fileList
}

// DeleteMedia will actually perform the deletion.
func DeleteMedia(delete bool, files []string) error {
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
