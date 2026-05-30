package models

type Job struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Payload    string `json:"payload"`
	Status     string `json:"status"`
	RetryCount int    `json:"retry_count"`
	Queue      string `json:"queue"`
}