package Education

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/reecerussell/ReeceRussellGo/Database"
)

// DataStore ... data access layer for education
type DataStore struct {
	Database Database.Database
}

// Init ... Initalizes data store
func (ds *DataStore) Init(db Database.Database) {
	ds.Database = db
}

// GetByID ... Gets education from table
func (ds *DataStore) GetByID(id int) (education Education, err error) {

	query := "SELECT TOP 1 * FROM [Education] WHERE [Id] = @Id AND [Hidden] = 0"

	rows, err := ds.Database.SelectByID(query, id)
	if err != nil {
		return Education{}, err
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&education.ID,
			&education.Title,
			&education.Organisation,
			&education.DateFrom,
			&education.DateTo,
			&education.Hidden)
		if err != nil {
			return Education{}, err
		}

		break
	}

	err = rows.Err()
	if err != nil {
		fmt.Println(err.Error())
		return Education{}, err
	}

	return education, nil
}

// Get ... Gets educations from table
func (ds *DataStore) Get() (educations []Education, err error) {
	query := "SELECT * FROM [dbo].[Education] WHERE [Hidden] = 0 ORDER BY [Id] ASC"

	rows, err := ds.Database.Select(query)
	if err != nil {
		return []Education{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var education Education

		err := rows.Scan(&education.ID,
			&education.Title,
			&education.Organisation,
			&education.DateFrom,
			&education.DateTo,
			&education.Hidden)
		if err != nil {
			return []Education{}, err
		}

		educations = append(educations, education)
	}
	err = rows.Err()
	if err != nil {
		return []Education{}, err
	}

	return educations, nil
}

// Add ... Insert education to table
func (ds *DataStore) Add(education Education) (id int, err error) {

	query := "INSERT INTO [Education] ([Title],[Organisation],[DateFrom],[DateTo],[Hidden]) VALUES (@Title,@Org,@From,@To,@Hidden)"

	lastID, err := ds.Database.Insert(query,
		sql.Named("Title", education.Title),
		sql.Named("Org", education.Organisation),
		sql.Named("From", education.DateFrom),
		sql.Named("To", education.DateTo),
		sql.Named("Hidden", education.Hidden))
	if err != nil {
		return 0, err
	}

	return int(lastID), nil
}

// Update ... Update education in table
func (ds *DataStore) Update(education Education) (err error) {
	if education.ID < 1 {
		return errors.New("Object has no ID")
	}

	query := "UPDATE [Education] SET [Title] = @Title, [Organisation] = @Org, [DateFrom] = @From, [DateTo] = @To, [Hidden] = @Hidden WHERE [Id] = @Id"

	rowCnt, err := ds.Database.Update(query,
		sql.Named("Title", education.Title),
		sql.Named("Org", education.Organisation),
		sql.Named("From", education.DateFrom),
		sql.Named("To", education.DateTo),
		sql.Named("Hidden", education.Hidden),
		sql.Named("Id", education.ID))
	if err != nil {
		return err
	}

	if rowCnt < 1 {
		return errors.New("No rows affected")
	}

	return nil
}

// Delete ... Delete education from table
func (ds *DataStore) Delete(id int) (err error) {
	if id < 1 {
		return errors.New("Invalid ID")
	}

	query := "DELETE FROM [Education] WHERE [Id] = @Id"

	rowCnt, err := ds.Database.Delete(query, sql.Named("Id", id))
	if err != nil {
		return err
	}

	if rowCnt < 1 {
		return errors.New("No rows affected")
	}

	return nil
}
