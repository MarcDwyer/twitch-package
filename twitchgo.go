// Package twitchgo fetches data from the Twitch API
package twitchgo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Twitch Create a new instance of your Twitch API
type Twitch struct {
	ClientID string
}

func (t Twitch) fetchTwitch(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.twitchtv.v5+json")
	req.Header.Set("Client-ID", t.ClientID)

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

//NewTwitchInstance Takes your Twitch client-ID and returns Twitch pointer
func NewTwitchInstance(id string) *Twitch {
	return &Twitch{
		ClientID: id,
	}
}

//GetTotalStreams Returns pointer to the total amount of livestreams on Twitch
func (t *Twitch) GetTotalStreams(language string) (*int, error) {
	url := fmt.Sprintf("https://api.twitch.tv/kraken/streams/?limit=1&language=%s", language)
	total, err := t.fetchTwitch(url)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving total")
	}
	var resp TResponse
	json.Unmarshal(total, &resp)
	return &resp.Total, nil
}

// GetStream returns a single stream
func (t *Twitch) GetStream(id int) (*SingleResponse, error) {
	url := fmt.Sprintf("https://api.twitch.tv/kraken/streams/%v", id)
	streamer, err := t.fetchTwitch(url)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving single stream")
	}
	var resp SingleResponse
	json.Unmarshal(streamer, &resp)
	return &resp, nil
}

// GetStreamList returns an array of streamers given a language, limit, and offset value
func (t *Twitch) GetStreamList(language string, limit *int, offset *int) (*TResponse, error) {
	url := fmt.Sprintf("https://api.twitch.tv/kraken/streams/?limit=%v&offset=%v&language=%s", *limit, *offset, language)
	list, err := t.fetchTwitch(url)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving list of streamers")
	}
	var resp TResponse
	json.Unmarshal(list, &resp)
	return &resp, nil
}
