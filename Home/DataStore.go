package Home

import (
	"context"

	"github.com/reecerussell/ReeceRussellGo/Education"
	"github.com/reecerussell/ReeceRussellGo/Experience"
	"github.com/reecerussell/ReeceRussellGo/Project"

	"github.com/reecerussell/ReeceRussellGo/Database"
	"github.com/reecerussell/ReeceRussellGo/Skills"
)

// DataStore ... Home datastore
type DataStore struct {
	Database Database.Database
}

// Init ... initialises data store
func (ds *DataStore) Init(database Database.Database) {
	ds.Database = database
}

// GetViewData ... main method to get all data for home page
func (ds *DataStore) GetViewData() (viewData ViewData) {
	ds.Database.Open()
	defer ds.Database.Close()

	viewData.Skills = ds.GetSkills()
	viewData.Projects = ds.GetProjects()
	viewData.Experience = ds.GetExperience()
	viewData.Education = ds.GetEducation()

	return viewData
}

// GetSkills ... get skills
func (ds *DataStore) GetSkills() (collection []Skills.Skills) {
	ctx := context.Background()

	if ds.Database.Db == nil {
		ds.Database.Open()
		// defer close ... assume this is being called elsewhere
		// and connection is no longer required.
		defer ds.Database.Close()
	} else {
		if err := ds.Database.Db.PingContext(ctx); err != nil {
			ds.Database.Open()
			// defer close ... assume this is being called elsewhere
			// and connection is no longer required.
			defer ds.Database.Close()
		}
	}

	skillTypesQuery := "SELECT * FROM [SkillTypes] WHERE [Hidden] = 0 ORDER BY [Id] ASC"
	skillsQuery := "SELECT * FROM [Skills] WHERE [Hidden] = 0 ORDER BY [Id] ASC"

	db := ds.Database.Db

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		return nil
	}

	// query roes for all skills
	rows, err := db.QueryContext(ctx, skillsQuery)
	if err != nil {
		return nil
	}

	defer rows.Close()

	var skills []Skills.Skill

	for rows.Next() {
		var skill Skills.Skill

		err := rows.Scan(&skill.ID,
			&skill.Skill,
			&skill.Type,
			&skill.Hidden)
		if err != nil {
			return nil
		}

		skills = append(skills, skill)
	}

	err = rows.Err()
	if err != nil {
		return nil
	}

	rows.Close()

	rows, err = db.QueryContext(ctx, skillTypesQuery)
	if err != nil {
		return nil
	}

	defer rows.Close()

	for rows.Next() {
		var (
			id     int
			name   string
			hidden bool
		)

		err := rows.Scan(&id, &name, &hidden)
		if err != nil {
			return nil
		}

		skillCollection := Skills.Skills{
			Title:  name,
			Skills: []Skills.Skill{},
		}

		for _, s := range skills {
			if s.Type == id {
				skillCollection.Skills = append(skillCollection.Skills, s)
			}
		}

		collection = append(collection, skillCollection)
	}

	err = rows.Err()
	if err != nil {
		return nil
	}

	return collection
}

// GetProjects ... get projects
func (ds *DataStore) GetProjects() (projects []Project.Project) {
	ctx := context.Background()

	if ds.Database.Db == nil {
		ds.Database.Open()
		// defer close ... assume this is being called elsewhere
		// and connection is no longer required.
		defer ds.Database.Close()
	} else {
		if err := ds.Database.Db.PingContext(ctx); err != nil {
			ds.Database.Open()
			// defer close ... assume this is being called elsewhere
			// and connection is no longer required.
			defer ds.Database.Close()
		}
	}

	query := "SELECT * FROM [Projects] WHERE [Hidden] = 0 ORDER BY [Id] DESC"

	db := ds.Database.Db

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil
	}

	defer rows.Close()

	for rows.Next() {
		var project Project.Project

		err := rows.Scan(
			&project.ID,
			&project.Name,
			&project.Description,
			&project.GithubLink,
			&project.ImageURL,
			&project.Teaser,
			&project.Hidden,
		)
		if err != nil {
			return nil
		}

		projects = append(projects, project)
	}

	err = rows.Err()
	if err != nil {
		return nil
	}

	return projects
}

// GetExperience ... get work
func (ds *DataStore) GetExperience() (experiences []Experience.Experience) {
	ctx := context.Background()

	if ds.Database.Db == nil {
		ds.Database.Open()
		// defer close ... assume this is being called elsewhere
		// and connection is no longer required.
		defer ds.Database.Close()
	} else {
		if err := ds.Database.Db.PingContext(ctx); err != nil {
			ds.Database.Open()
			// defer close ... assume this is being called elsewhere
			// and connection is no longer required.
			defer ds.Database.Close()
		}
	}

	query := "SELECT * FROM [Experience] WHERE [Hidden] = 0 ORDER BY [Id] DESC"

	db := ds.Database.Db

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil
	}

	defer rows.Close()

	for rows.Next() {
		var experience Experience.Experience

		err := rows.Scan(
			&experience.ID,
			&experience.Title,
			&experience.Description,
			&experience.Organisation,
			&experience.DateFrom,
			&experience.DateTo,
			&experience.Hidden,
		)
		if err != nil {
			return nil
		}

		experiences = append(experiences, experience)
	}

	err = rows.Err()
	if err != nil {
		return nil
	}

	return experiences
}

// GetEducation ... get education
func (ds *DataStore) GetEducation() (educations []Education.Education) {
	ctx := context.Background()

	if ds.Database.Db == nil {
		ds.Database.Open()
		// defer close ... assume this is being called elsewhere
		// and connection is no longer required.
		defer ds.Database.Close()
	} else {
		if err := ds.Database.Db.PingContext(ctx); err != nil {
			ds.Database.Open()
			// defer close ... assume this is being called elsewhere
			// and connection is no longer required.
			defer ds.Database.Close()
		}
	}

	query := "SELECT * FROM [Education] WHERE [Hidden] = 0 ORDER BY [Id] DESC"

	db := ds.Database.Db

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil
	}

	defer rows.Close()

	for rows.Next() {
		var education Education.Education

		err := rows.Scan(
			&education.ID,
			&education.Title,
			&education.Description,
			&education.Organisation,
			&education.DateFrom,
			&education.DateTo,
			&education.Hidden,
		)
		if err != nil {
			return nil
		}

		educations = append(educations, education)
	}

	err = rows.Err()
	if err != nil {
		return nil
	}

	return educations
}
