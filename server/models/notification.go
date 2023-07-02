package models

type Notification struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}
