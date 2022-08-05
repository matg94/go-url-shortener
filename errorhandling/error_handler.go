package errorhandling

import (
	"encoding/json"
	"log"
)

type AppError struct {
	Err      string `json:"error"`
	Location string `json:"location"`
	Cause    string `json:"cause"`
}

func (err *AppError) ToJSON() string {
	error_json, _ := json.Marshal(err)
	return string(error_json)
}

func CreateError(err error, location, values string) *AppError {
	return &AppError{
		Err:      err.Error(),
		Location: location,
		Cause:    values,
	}
}

func HandleError(err error, location, values string) {
	if err == nil {
		return
	}
	appError := CreateError(err, location, values).ToJSON()
	log.Print(appError)
}
