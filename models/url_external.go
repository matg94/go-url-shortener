package models

import (
	"encoding/json"
	"errors"

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
	if err != nil {
		errorhandling.HandleError(err, "Shorten request parsing", string(json_data))
		return request, err
	}
	if request.URL == "" { // TODO: Add proper errors
		errorhandling.HandleError(errors.New("URL not defined in request"), "Shorten request parsing", request.URL)
		return request, errors.New("URL not defined in request")
	}
	return request, err
}

func LongRequestFromJson(json_data []byte) (URLElongateResponse, error) {
	request := URLElongateResponse{}
	err := json.Unmarshal(json_data, &request)
	if err != nil {
		errorhandling.HandleError(err, "Elongate request parsing", string(json_data))
		return request, err
	}
	if request.Hash == "" { // TODO: Add proper errors
		errorhandling.HandleError(errors.New("URL not defined in request"), "Elongate request parsing", request.Hash)
		return request, errors.New("URL not defined in request")
	}
	return request, err
}
