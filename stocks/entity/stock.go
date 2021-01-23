package entity

// Stock is the entity to be stored and returned in JSON
type Stock struct {
	ID     uint   `json:"id" binding:required`
	Name   string `json:"name" binding:"required"`
	Ticker string `json:"ticker" binding:"required"`
}
