package models

// Response is JSON view model for all controller responses
type Response struct {
	Success   bool        `json:"success"`
	RequestID string      `json:"request_id"`
	Data      interface{} `json:"data"`
	Error     interface{} `json:"error"`
}
