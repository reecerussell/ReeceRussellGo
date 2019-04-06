package Experience

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/reecerussell/ReeceRussellGo/Database"
)

type ExperienceDataStore struct {
	Database Database.Database
}

// Init ... Initalizes data store
func (ds *ExperienceDataStore) Init(db Database.Database) {
	ds.Database = db
}

// GetByID ... Gets project from table
func (ds *ExperienceDataStore) GetByID(id int) (experience Experience, err error) {

	query := "SELECT TOP 1 * FROM [Experience] WHERE [Id] = @Id AND [Hidden] = 0"

	rows, err := ds.Database.SelectByID(query, id)
	if err != nil {
		return Experience{}, err
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&experience.ID,
			&experience.Title,
			&experience.Description,
			&experience.Organisation,
			&experience.DateFrom,
			&experience.DateTo,
			&experience.Hidden)
		if err != nil {
			return Experience{}, err
		}

		break
	}

	err = rows.Err()
	if err != nil {
		fmt.Println(err.Error())
		return Experience{}, err
	}

	return experience, nil
}

// Get ... Gets projects from table
func (ds *ExperienceDataStore) Get() (experiences []Experience, err error) {
	query := "SELECT * FROM [dbo].[Experience] WHERE [Hidden] = 0 ORDER BY [Id] ASC"

	rows, err := ds.Database.Select(query)
	if err != nil {
		return []Experience{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var experience Experience

		err := rows.Scan(&experience.ID,
			&experience.Title,
			&experience.Description,
			&experience.Organisation,
			&experience.DateFrom,
			&experience.DateTo,
			&experience.Hidden)
		if err != nil {
			return []Experience{}, err
		}

		experiences = append(experiences, experience)
	}
	err = rows.Err()
	if err != nil {
		return []Experience{}, err
	}

	return experiences, nil
}

// Add ... Insert experience to table
func (ds *ExperienceDataStore) Add(experience Experience) (id int, err error) {

	query := "INSERT INTO [Experience] ([Title],[Description],[Organisation],[DateFrom],[DateTo],[Hidden]) VALUES (@Title,@Desc,@Org,@From,@To,@Hidden)"

	lastID, err := ds.Database.Insert(query,
		sql.Named("Title", experience.Title),
		sql.Named("Desc", experience.Description),
		sql.Named("Org", experience.Organisation),
		sql.Named("From", experience.DateFrom),
		sql.Named("To", experience.DateTo),
		sql.Named("Hidden", experience.Hidden))
	if err != nil {
		return 0, err
	}

	return int(lastID), nil
}

// Update ... Update experience in table
func (ds *ExperienceDataStore) Update(experience Experience) (err error) {
	if experience.ID < 1 {
		return errors.New("Object has no ID")
	}

	query := "UPDATE [Experience] SET [Title] = @Title, [Description] = @Desc, [Organisation] = @Org, [DateFrom] = @From, [DateTo] = @To, [Hidden] = @Hidden WHERE [Id] = @Id"

	rowCnt, err := ds.Database.Update(query,
		sql.Named("Title", experience.Title),
		sql.Named("Desc", experience.Description),
		sql.Named("Org", experience.Organisation),
		sql.Named("From", experience.DateFrom),
		sql.Named("To", experience.DateTo),
		sql.Named("Hidden", experience.Hidden),
		sql.Named("Id", experience.ID))
	if err != nil {
		return err
	}

	if rowCnt < 1 {
		return errors.New("No rows affected")
	}

	return nil
}

// Delete ... Delete experience from table
func (ds *ExperienceDataStore) Delete(id int) (err error) {
	if id < 1 {
		return errors.New("Invalid ID")
	}

	query := "DELETE FROM [experience] WHERE [Id] = @Id"

	rowCnt, err := ds.Database.Delete(query, sql.Named("Id", id))
	if err != nil {
		return err
	}

	if rowCnt < 1 {
		return errors.New("No rows affected")
	}

	return nil
}
