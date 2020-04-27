package models

// ResponseError view model
type ResponseError struct {
	Code       int               `json:"code"`
	Cause      string            `json:"cause"`
	Message    string            `json:"message"`
	Stack      string            `json:"stack"`
	Validation map[string]string `json:"validation"`
}
