package locator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"git.linuxrocker.com/mattburchett/Housekeeper/pkg/config"
	"git.linuxrocker.com/mattburchett/Housekeeper/pkg/model"
)

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

func getTitles(config config.Config, sectionID int) string {

}
