package fees

type FeeItem struct {
	Name             string  `json:"name"`
	Rate             float64 `json:"rate"`
	HousekeepingNote string  `json:"housekeeping_note"`
}
