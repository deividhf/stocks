package entity

// Stock is the entity to be stored and returned in JSON
type Stock struct {
	Name   string `json:"name" binding:"required"`
	Ticker string `json:"ticker" binding:"required"`
}
