package models

type Response struct {
	Status int         `json:"status"`
	Body   interface{} `json:"body,omitempty"`
}
