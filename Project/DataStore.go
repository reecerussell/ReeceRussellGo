package Project

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/reecerussell/ReeceRussellGo/Database"
)

// ProjectDataStore ... Datastore object
type ProjectDataStore struct {
	Database Database.Database
}

// Init ... Initalizes data store
func (ds *ProjectDataStore) Init(db Database.Database) {
	ds.Database = db
}

// GetByID ... Gets project from table
func (ds *ProjectDataStore) GetByID(id int) (project Project, err error) {

	query := "SELECT TOP 1 * FROM [Projects] WHERE [Id] = @Id AND [Hidden] = 0"

	rows, err := ds.Database.SelectByID(query, id)
	if err != nil {
		return Project{}, err
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&project.ID,
			&project.Name,
			&project.Description,
			&project.GithubLink,
			&project.ImageURL,
			&project.Teaser,
			&project.Hidden)
		if err != nil {
			fmt.Println(err.Error())
			return Project{}, err
		}

		break
	}

	err = rows.Err()
	if err != nil {
		fmt.Println(err.Error())
		return Project{}, err
	}

	return project, nil
}

// Get ... Gets projects from table
func (ds *ProjectDataStore) Get() (projects []Project, err error) {
	query := "SELECT * FROM [Projects] WHERE [Hidden] = 0 ORDER BY [Id] ASC"

	rows, err := ds.Database.Select(query)
	if err != nil {
		return []Project{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var project Project

		err := rows.Scan(&project.ID,
			&project.Name,
			&project.Description,
			&project.GithubLink,
			&project.ImageURL,
			&project.Teaser,
			&project.Hidden)
		if err != nil {
			return []Project{}, err
		}

		projects = append(projects, project)
	}
	err = rows.Err()
	if err != nil {
		return []Project{}, err
	}

	return projects, nil
}

// Add ... Insert project to table
func (ds *ProjectDataStore) Add(project Project) (id int, err error) {

	query := "INSERT INTO [Projects] ([Name],[Description],[GithubLink],[ImageUrl],[Teaser],[Hidden]) VALUES (@Name,@Desc,@Git,@Img,@Teaser,@Hidden)"

	rowsAffected, lastID, err := ds.Database.Insert(query,
		sql.Named("Name", project.Name),
		sql.Named("Desc", project.Description),
		sql.Named("Git", project.GithubLink),
		sql.Named("Img", project.ImageURL),
		sql.Named("Teaser", project.Teaser),
		sql.Named("Hidden", project.Hidden))
	if err != nil {
		return 0, err
	}

	if rowsAffected < 1 {
		return 0, errors.New("No rows affected")
	}

	return int(lastID), nil
}

// Update ... Update project in table
func (ds *ProjectDataStore) Update(project Project) (err error) {
	if project.ID < 1 {
		return errors.New("Project has no ID")
	}

	query := "UPDATE [Projects] SET [Name] = @Name, [Description] = @Desc, [GithubLink] = @Git, [ImageUrl] = @Img, [Teaser] = @Teaser, [Hidden] = @Hidden WHERE [Id] = @Id"

	rowCnt, err := ds.Database.Update(query,
		sql.Named("Name", project.Name),
		sql.Named("Desc", project.Description),
		sql.Named("Git", project.GithubLink),
		sql.Named("Img", project.ImageURL),
		sql.Named("Teaser", project.Teaser),
		sql.Named("Hidden", project.Hidden),
		sql.Named("Id", project.ID))
	if err != nil {
		return err
	}

	if rowCnt < 1 {
		return errors.New("No rows affected")
	}

	return nil
}

// Delete ... Delete project from table
func (ds *ProjectDataStore) Delete(id int) (err error) {
	if id < 1 {
		return errors.New("Invalid project ID")
	}

	query := "DELETE FROM [Projects] WHERE [Id] = @Id"

	rowCnt, err := ds.Database.Delete(query, sql.Named("Id", id))
	if err != nil {
		return err
	}

	if rowCnt < 1 {
		return errors.New("No rows affected")
	}

	return nil
}
