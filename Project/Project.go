package Project

import "github.com/reecerussell/ReeceRussellGo/Nullables"

// Project ... Project object
type Project struct {
	ID          int                  `json:"id"`
	Name        Nullables.NullString `json:"name"`
	Description Nullables.NullString `json:"description"`
	GithubLink  Nullables.NullString `json:"githubLink"`
	ImageURL    Nullables.NullString `json:"imageUrl"`
	Teaser      Nullables.NullString `json:"teaser"`
	Hidden      bool                 `json:"hidden"`
}
