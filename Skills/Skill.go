package Skills

import (
	"github.com/reecerussell/ReeceRussellGo/Nullables"
)

// Skill object
type Skill struct {
	ID     int                  `json:"id"`
	Skill  Nullables.NullString `json:"skill"`
	Type   int                  `json:"type"`
	Hidden bool                 `json:"hidden"`
}

// Skills ... a collection of skill with their linked title
type Skills struct {
	Title  string  `json:"title"`
	Skills []Skill `json:"skills"`
}
