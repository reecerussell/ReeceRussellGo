package Experience

// Experience ... experience object
type Experience struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Organisation string `json:"organisation"`
	DateFrom     string `json:"from"`
	DateTo       string `json:"to"`
	Hidden       bool   `json:"hidden"`
}
