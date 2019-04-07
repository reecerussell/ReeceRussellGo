package Home

import (
	"github.com/reecerussell/ReeceRussellGo/Education"
	"github.com/reecerussell/ReeceRussellGo/Experience"
	"github.com/reecerussell/ReeceRussellGo/Project"
	"github.com/reecerussell/ReeceRussellGo/Skills"
)

// ViewData ... model for home view data
type ViewData struct {
	Projects   []Project.Project       `json:"projects"`
	Experience []Experience.Experience `json:"experience"`
	Education  []Education.Education   `json:"education"`
	Skills     []Skills.Skills         `json:"skills"`
}
