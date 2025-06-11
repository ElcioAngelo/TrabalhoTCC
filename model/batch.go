package model

import "time"

type Batch struct {
	ID int 						`json:"id"`
	Code string					`json:"code"`
	ExpirationDate time.Time	`json:"expiration_date"`
}

