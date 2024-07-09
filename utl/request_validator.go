package utl

import (
	"errors"
	"strings"

	"github.com/go-playground/locales/en_US"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

const (
	requestValidationErrorTypeInvalidField RequestValidationErrorType = "invalid_field"
)

type RequestValidationErrorType string

type RequestValidationErrorResponse struct {
	ErrorType RequestValidationErrorType `json:"error_type"`
	Errors    []RequestValidationError   `json:"errors"`
}

type RequestValidationError struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Message string `json:"message"`
}

func ValidateRequestAllEmptyResponse() RequestValidationErrorResponse {
	return RequestValidationErrorResponse{
		ErrorType: requestValidationErrorTypeInvalidField,
		Errors: []RequestValidationError{
			{
				Field:   "all",
				Message: "all fields are required",
			},
		},
	}
}

func ValidateRequestStruct(req interface{}) (RequestValidationErrorResponse, bool) {
	v := validator.New()

	locale := en_US.New()
	uni := ut.New(locale, locale)
	trans, _ := uni.GetTranslator(locale.Locale())

	err := v.Struct(req)
	if err == nil {
		return RequestValidationErrorResponse{}, true
	}

	var validationErrors validator.ValidationErrors
	errors.As(err, &validationErrors)

	var apiValidationErrors []RequestValidationError
	for _, validationError := range validationErrors {
		apiValidationErrors = append(apiValidationErrors, RequestValidationError{
			Field:   strings.ToLower(validationError.Field()),
			Tag:     strings.ToLower(validationError.Tag()),
			Message: validationError.Translate(trans),
		})
	}

	if apiValidationErrors != nil {
		return RequestValidationErrorResponse{
			ErrorType: requestValidationErrorTypeInvalidField,
			Errors:    apiValidationErrors,
		}, false
	}

	return RequestValidationErrorResponse{}, true
}
