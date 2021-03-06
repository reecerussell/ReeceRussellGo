package Database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	// ... driver for database/sql
	_ "github.com/denisenkom/go-mssqldb"
)

var (
	server           = "54.38.215.159"
	port             = "11433"
	database         = "ReeceRussell_Backend"
	user             = "sa"
	password         = "#Exchange_2013!"
	connectionString = fmt.Sprintf("server=%s;port=%s;database=%s;user id=%s; password=%s;", server, port, database, user, password)
)

// Database ... Database type
type Database struct {
	Db *sql.DB
}

// Open ... Opens connection to database
func (database *Database) Open() {
	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	database.Db = db
}

// Close ... Closes connection to the database
func (database *Database) Close() {
	err := database.Db.Close()
	if err != nil {
		fmt.Println("Failed to close connection to database")
	}
}

// Select ... Executes select command
func (database *Database) Select(query string) (rows *sql.Rows, err error) {

	database.Open()
	defer database.Close()

	db := database.Db

	ctx := context.Background()

	// Check if database is alive.
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	// Execute query
	rows, err = db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

// SelectByID ... Selects row by id
func (database *Database) SelectByID(query string, id int) (rows *sql.Rows, err error) {

	database.Open()
	defer database.Close()

	db := database.Db

	ctx := context.Background()

	// Check if database is alive.
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	// Execute query
	rows, err = db.QueryContext(ctx, query, sql.Named("Id", id))
	if err != nil {
		return nil, err
	}

	return rows, nil
}

// Insert ... Inserts record into database
func (database *Database) Insert(query string, params ...sql.NamedArg) (newID int64, err error) {

	database.Open()
	defer database.Close()

	db := database.Db

	ctx := context.Background()

	if db == nil {
		err = errors.New("db is null")
		return -1, err
	}

	// Check if database is alive.
	err = db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	identitySQL := " SELECT SCOPE_IDENTITY()"
	query = query + identitySQL

	stmt, err := db.Prepare(query)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	interfaceParams := make([]interface{}, len(params))
	for i, d := range params {
		interfaceParams[i] = d
	}

	row := stmt.QueryRowContext(ctx, interfaceParams...)

	err = row.Scan(&newID)
	if err != nil {
		return -1, err
	}

	return newID, nil
}

// Update ... Provides sutible method for updating records
func (database *Database) Update(query string, params ...sql.NamedArg) (rowCnt int64, err error) {

	database.Open()
	defer database.Close()

	db := database.Db

	ctx := context.Background()

	// Check if database is alive.
	err = db.PingContext(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return -1, err
	}

	interfaceParams := make([]interface{}, len(params))
	for i, d := range params {
		interfaceParams[i] = d
	}

	// Execute non-query with named parameters
	result, err := db.ExecContext(
		ctx,
		query,
		interfaceParams...)
	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}

// Delete ... Provides delete method
func (database *Database) Delete(query string, params ...sql.NamedArg) (rowCnt int64, err error) {
	return database.Update(query, params...)
}
