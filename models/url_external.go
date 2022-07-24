package models

import "encoding/json"

type URLShortenRequest struct {
	URL string `json:"url"`
}

type URLShortenResponse struct {
	Hash string `json:"url"`
}

type URLElongateResponse struct {
	URL string `json:"url"`
}

func ShortenRequestFromJson(json_data string) URLShortenRequest {
	request := URLShortenRequest{}
	json.Unmarshal([]byte(json_data), &request)
	return request
}
