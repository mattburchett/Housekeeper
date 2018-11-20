package locator

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"

	"git.linuxrocker.com/mattburchett/Housekeeper/pkg/config"
	"git.linuxrocker.com/mattburchett/Housekeeper/pkg/model"
	"git.linuxrocker.com/mattburchett/Housekeeper/pkg/util"
)

// GetLibraryType checks to see what type the library is.
func GetLibraryType(config config.Config, sectionID int) string {
	typeURL := fmt.Sprintf("%s:%d%s%d%s%s", config.PlexHost, config.PlexPort, "/library/sections/", sectionID, "/?X-Plex-Token=", config.PlexToken)

	req, err := http.NewRequest(http.MethodGet, typeURL, nil)
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

	typeModel := model.XMLPlexLibraryType{}
	xml.Unmarshal(body, &typeModel)

	var libraryType string

	if typeModel.Thumb == "/:/resources/movie.png" {
		libraryType = "movie"
	} else if typeModel.Thumb == "/:/resources/show.png" {
		libraryType = "show"
	} else {
		log.Fatal("Unsupported library type found. This app only supports movies and shows.")
	}

	return libraryType
}

// GetCount will gather a count of media in a specific library, required for GetTitles.
func GetCount(config config.Config, sectionID int) int {
	countURL := fmt.Sprintf("%s%s%s%s%s%d", config.BaseURL, config.PlexPyContext, "/api/v2?apikey=", config.PlexPyAPIKey, "&cmd=get_library&section_id=", sectionID)
	req, err := http.NewRequest(http.MethodGet, countURL, nil)
	if err != nil {
		log.Fatal(err)
	}

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

	countModel := model.PlexPyLibraryInfo{}
	jsonErr := json.Unmarshal(body, &countModel)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	count := countModel.Response.Data.Count

	return count
}

// GetTitles will gather a list of information from the media in the library, based on the previous count.
func GetTitles(config config.Config, sectionID int, days int) ([]int, []string) {
	count := GetCount(config, sectionID)

	titlesURL := fmt.Sprintf("%s%s%s%s%s%d%s%d", config.BaseURL, config.PlexPyContext, "/api/v2?apikey=", config.PlexPyAPIKey, "&cmd=get_library_media_info&section_id=", sectionID, "&order_column=last_played&refresh=true&order_dir=asc&length=", count)

	req, err := http.NewRequest(http.MethodGet, titlesURL, nil)
	if err != nil {
		log.Fatal(err)
	}

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

	titleModel := model.PlexPyMediaInfo{}
	jsonErr := json.Unmarshal(body, &titleModel)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	data := titleModel.Response.Data.Data
	titles := make([]string, 0)
	ids := make([]int, 0)

	epoch := util.SubtractedEpoch(days)

	for _, i := range data {
		if int64(i.LastPlayed) <= epoch && int64(i.LastPlayed) != 0 {
			titles = append(titles, i.Title)
			strirk, err := strconv.Atoi(i.RatingKey)
			if err != nil {
				log.Fatal(err)
			}
			ids = append(ids, strirk)
		}
		if i.LastPlayed <= 0 {
			if i.AddedAt <= epoch {
				titles = append(titles, i.Title)
				strirk, err := strconv.Atoi(i.RatingKey)
				if err != nil {
					log.Fatal(err)
				}
				ids = append(ids, strirk)
			}
		}
	}
	sort.Strings(titles)
	return ids, titles
}
