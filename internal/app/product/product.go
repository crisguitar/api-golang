package product

type Product struct {
	Id         string  `json:"developmentId" db:"development_id"`
	Colour     string `json:"colour" db:"colour"`
	FabricType string `json:"fabricType" db:"fabric_type"`
	Season     string `json:"season" db:"season"`
	Status     string `json:"status" db:"status"`
}
