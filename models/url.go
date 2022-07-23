package models

import (
	"encoding/json"
)

type URL struct {
	LongURL string `json:"LongURL"`
	Hits    uint   `json:"Hits"`
}

func (url *URL) ToJSON() (string, error) {
	url_json, err := json.Marshal(url)
	return string(url_json), err
}

func FromJSON(json_data string) (URL, error) {
	var url URL
	err := json.Unmarshal([]byte(json_data), &url)
	return url, err
}
