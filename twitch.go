// Package twitch fetches data from the Twitch API
package twitch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//Twitch Create a new instance of your Twitch API
type Twitch struct {
	clientID string
}

func fetchTwitch(url string, method string) ([]byte, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.twitchtv.v5+json")
	req.Header.Set("Client-ID", os.Getenv("TWITCH"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}

//SetTwitchID Takes your Twitch client-ID and returns Twitch pointer
func SetTwitchID(id string) *Twitch {
	return &Twitch{
		clientID: id,
	}
}

//GetTotalStreams Returns pointer to total amount of livestreams on Twitch
func (t *Twitch) GetTotalStreams(language string) (*int, error) {
	url := fmt.Sprintf("https://api.twitch.tv/kraken/streams/?limit=1&language=%s", language)
	total, err := fetchTwitch(url, "GET")
	if err != nil {
		return nil, fmt.Errorf("Error retrieving total")
	}
	var resp TResponse
	json.Unmarshal(total, &resp)
	return &resp.Total, nil
}
