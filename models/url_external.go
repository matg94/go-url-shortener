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

func ShortenRequestFromJson(json_data []byte) URLShortenRequest {
	request := URLShortenRequest{}
	json.Unmarshal(json_data, &request)
	return request
}

func LongRequestFromJson(json_data []byte) URLElongateResponse {
	request := URLElongateResponse{}
	json.Unmarshal(json_data, &request)
	return request
}
