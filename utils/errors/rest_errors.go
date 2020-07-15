package errors

import "net/http"

type RestErr struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
	Error   string `json:"Error"`
}

func NewBadRequestError(message string) *RestErr {

	saveErr := RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
	return &saveErr
}

func NewNotFoundError(message string) *RestErr {

	saveErr := RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
	return &saveErr
}

func NewInternalServerError(message string) *RestErr {

	saveErr := RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
	return &saveErr
}
