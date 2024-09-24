package models

import (
	"time"
)

// data type for each point addition transaction requiring all fields
type Transaction struct {
	Points     int    `json:"points" binding:"required"`
	Payer      string `json:"payer" binding:"required"`
	Timestamp  string `json:"timestamp" bindgind:"required"`
	ParsedTime time.Time
}

// data type for spending request
type SpendRequest struct {
	Points int
}
