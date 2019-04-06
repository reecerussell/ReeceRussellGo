package Project

// Project ... Project object
type Project struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	GithubLink  string `json:"githubLink"`
	ImageURL    string `json:"imageUrl"`
	Teaser      string `json:"teaser"`
	Hidden      bool   `json:"hidden"`
}
