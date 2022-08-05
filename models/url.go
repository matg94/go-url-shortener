package models

import (
	"encoding/json"

	"github.com/matg94/go-url-shortener/errorhandling"
)

type URL struct {
	LongURL string `json:"LongURL"`
	Hits    uint   `json:"Hits"`
}

func (url *URL) ToJSON() (string, error) {
	url_json, err := json.Marshal(url)
	errorhandling.HandleError(err, "URL object To JSON", url.LongURL)
	return string(url_json), err
}

func FromJSON(json_data string) (URL, error) {
	var url URL
	err := json.Unmarshal([]byte(json_data), &url)
	errorhandling.HandleError(err, "JSON To URL object", json_data)
	return url, err
}
