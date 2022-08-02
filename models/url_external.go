package models

import (
	"encoding/json"
)

type URLShortenRequest struct {
	URL string `json:"url"`
}

type URLShortenResponse struct {
	Hash string `json:"url"`
}

type URLElongateResponse struct {
	URL string `json:"url"`
}

func ShortenRequestFromJson(json_data []byte) (URLShortenRequest, error) {
	request := URLShortenRequest{}
	err := json.Unmarshal(json_data, &request)
	return request, err
}

func LongRequestFromJson(json_data []byte) (URLElongateResponse, error) {
	request := URLElongateResponse{}
	err := json.Unmarshal(json_data, &request)
	return request, err
}
