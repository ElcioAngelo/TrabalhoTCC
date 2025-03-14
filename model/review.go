package model

type Review struct {
	ID int				`json:"id"`
	Description string	`json:"description"`
	Rating		string  `json:"rating"`
}

