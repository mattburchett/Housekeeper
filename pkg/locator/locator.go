package locator

import (
	"encoding/json"
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
		if i.LastPlayed < epoch {
			titles = append(titles, i.Title)
			strirk, err := strconv.Atoi(i.RatingKey)
			if err != nil {
				log.Fatal(err)
			}
			ids = append(ids, strirk)
		}
		if i.LastPlayed < 0 {
			stri, err := strconv.Atoi(i.AddedAt)
			if err != nil {
				log.Fatal(err)
			}
			if int64(stri) < epoch {
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
