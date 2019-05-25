package database

import (
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

func TestFindDatabaseRecord(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// our sql.DB must exec the follow query
	mock.ExpectQuery("SELECT \\* FROM my_table").
		// the number 3 must be on the query like "where id = 3"
		WithArgs(3)
	// TODO: define the values that need to be returned from the database

	myDB := NewRepository(db) // passes the mock to our code

	if err = myDB.Find(3); err != nil {
		t.Errorf("Something went wrong: %s", err.Error())
	}

}
