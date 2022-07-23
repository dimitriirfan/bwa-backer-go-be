package helper

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func FormatError(err error) []string {

	var finalErrors []string
	var ve validator.ValidationErrors

	if errors.As(err, &ve) {
		for _, e := range err.(validator.ValidationErrors) {
			finalErrors = append(finalErrors, e.Error())
		}
	} else {
		finalErrors = append(finalErrors, err.Error())

	}

	return finalErrors
}
