package Experience

import "github.com/reecerussell/ReeceRussellGo/Nullables"

// Experience ... experience object
type Experience struct {
	ID           int                  `json:"id"`
	Title        Nullables.NullString `json:"title"`
	Description  Nullables.NullString `json:"description"`
	Organisation Nullables.NullString `json:"organisation"`
	DateFrom     Nullables.NullString `json:"from"`
	DateTo       Nullables.NullString `json:"to"`
	Hidden       bool                 `json:"hidden"`
}
