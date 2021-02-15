package lmpm

type Guest struct {
	Name     string   `json:"name"`
	Phone    string   `json:"phone"`
	Email    string   `json:"email"`
	Location Location `json:"location"`
}

type Location struct {
	Country    string `json:"country"`
	State      string `json:"state"`
	City       string `json:"city"`
	PostalCode string `json:"postal_code"`
}
