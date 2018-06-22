package product

import . "github.com/satori/go.uuid"

type Product struct {
	Id         UUID   `json:"developmentId" db:"development_id"`
	Colour     string `json:"colour" db:"colour"`
	FabricType string `json:"fabricType" db:"fabric_type"`
	Season     string `json:"season" db:"season"`
	Status     string `json:"status" db:"status"`
}
