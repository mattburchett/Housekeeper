package communicator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"git.linuxrocker.com/mattburchett/Housekeeper/pkg/config"
)

// TelegramPost will send a message to a specific ChatID in Telegram containing the list of items to be cleaned with this cleaner.
func TelegramPost(config config.Config, titles []string) error {
	url := "https://api.telegram.org/bot" + config.TelegramToken + "/sendMessage"

	values := map[string]string{"chat_id": config.TelegramChatID, "text": "The following items are to be removed from " + config.ServerName + " in 24 hours. Please go to Plex and start the title to keep it on" + config.ServerName + ". You do not need to keep watching, just hit play and load a few seconds.\n\n" + fmt.Sprintf("%v", strings.Join(titles, "\n")), "disable_notifications": "true"}

	jsonValue, _ := json.Marshal(values)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	req.Header.Set("X-Custom-Header", "Housekeeper")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return err
}
