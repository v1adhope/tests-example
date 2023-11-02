package entity

type Member struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       uint8  `json:"age"`
}
