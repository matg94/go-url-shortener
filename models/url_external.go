package models

import (
	"encoding/json"

	"github.com/matg94/go-url-shortener/errorhandling"
)

type URLShortenRequest struct {
	URL string `json:"url"`
}

type URLElongateResponse struct {
	Hash string `json:"url"`
}

func ShortenRequestFromJson(json_data []byte) (URLShortenRequest, error) {
	request := URLShortenRequest{}
	err := json.Unmarshal(json_data, &request)
	errorhandling.HandleError(err, "Shorten request parsing", string(json_data))
	return request, err
}

func LongRequestFromJson(json_data []byte) (URLElongateResponse, error) {
	request := URLElongateResponse{}
	err := json.Unmarshal(json_data, &request)
	errorhandling.HandleError(err, "Elongate request parsing", string(json_data))
	return request, err
}
