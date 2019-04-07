package Skills

import (
	"database/sql"
	"errors"

	"github.com/reecerussell/ReeceRussellGo/Database"
)

// DataStore ... Skill data store
type DataStore struct {
	Database Database.Database
}

// Init ... Initialises data store
func (ds *DataStore) Init(db Database.Database) {
	ds.Database = db
}

// GetByID ... Gets skill from table
func (ds *DataStore) GetByID(id int) (skill Skill, err error) {
	query := "SELECT TOP 1 * FROM [Skills] WHERE [Id] = @Id AND [Hidden] = 0"

	rows, err := ds.Database.SelectByID(query, id)
	if err != nil {
		return Skill{}, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&skill.ID,
			&skill.Skill,
			&skill.Type,
			&skill.Hidden)
		if err != nil {
			return Skill{}, err
		}

		break
	}

	err = rows.Err()
	if err != nil {
		return Skill{}, err
	}

	return skill, nil
}

// Get ... Gets skills from table
func (ds *DataStore) Get() (skills []Skill, err error) {
	query := "SELECT * FROM [Skills] WHERE [Hidden] = 0 ORDER BY [Id] ASC"

	rows, err := ds.Database.Select(query)
	if err != nil {
		return []Skill{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var skill Skill

		err := rows.Scan(&skill.ID,
			&skill.Skill,
			&skill.Type,
			&skill.Hidden)
		if err != nil {
			return []Skill{}, err
		}

		skills = append(skills, skill)
	}

	err = rows.Err()
	if err != nil {
		return []Skill{}, err
	}

	return skills, nil
}

// Add ... Adds skill to the table
func (ds *DataStore) Add(skill Skill) (id int, err error) {
	query := "INSERT INTO [Skills] ([Skill],[Type],[Hidden]) VALUES (@Skill,@Type,@Hidden)"

	lastID, err := ds.Database.Insert(query,
		sql.Named("Skill", skill.Skill),
		sql.Named("Type", skill.Type),
		sql.Named("Hidden", skill.Hidden))
	if err != nil {
		return 0, err
	}

	return int(lastID), nil
}

// Update ... Update a skill
func (ds *DataStore) Update(skill Skill) (err error) {
	if skill.ID < 1 {
		return errors.New("Skill has no ID")
	}

	query := "UPDATE [Skills] SET [Skill] = @Skill, [Type] = @Type, [Hidden] = @Hidden WHERE [Id] = @Id"

	rowCnt, err := ds.Database.Update(query,
		sql.Named("Skill", skill.Skill),
		sql.Named("Type", skill.Type),
		sql.Named("Hidden", skill.Hidden))
	if err != nil {
		return err
	}

	if rowCnt < 1 {
		return errors.New("No rows affected")
	}

	return nil
}

// Delete ... Deletes a skill
func (ds *DataStore) Delete(id int) (err error) {
	if id < 1 {
		return errors.New("Invalid skill id")
	}

	query := "DELETE FROM [Skills] WHERE [Id] = @Id"

	rowCnt, err := ds.Database.Delete(query, sql.Named("Id", id))
	if err != nil {
		return err
	}

	if rowCnt < 1 {
		return errors.New("No rows affected")
	}

	return nil
}
