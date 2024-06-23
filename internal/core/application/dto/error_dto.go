package dto

import "github.com/duducv/go-clean-arch/internal/core/application/constants"

type ErrorOutputLayer string

type ErrorOutput struct {
	Message    []string        `json:"message"`
	Raw        string          `json:"raw"`
	Layer      constants.Layer `json:"layer"`
	StatusCode int             `json:"statusCode"`
}

func NewErrorOutput(
	raw string,
	layer constants.Layer,
	statusCode int,
	message ...string,
) *ErrorOutput {
	return &ErrorOutput{
		Message:    message,
		Raw:        raw,
		Layer:      layer,
		StatusCode: statusCode,
	}
}
