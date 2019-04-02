package Home

import (
	"github.com/reecerussell/ReeceRussellGo/Experience"
	"github.com/reecerussell/ReeceRussellGo/Project"
)

// ViewData ... model for home view data
type ViewData struct {
	Projects   []Project.Project       `json:"projects"`
	Experience []Experience.Experience `json:"experience"`
}
