package database

import (
	"database/sql"
	"errors"
	"strconv"
)

//Database
type Repository interface {
	Find(id int) error
}

//A string type implement Repository interface
type repository struct {
	database *sql.DB
}

//method that we want to test
func (r repository) Find(id int) error {
	return errors.New("repo's Find method")
}

//________________________
//________________________

//Error
/*
type Error interface {
	Error() string
}
*/

type CustomError struct {
	message string
}

func (e CustomError) Error() string {
	return e.message
}

//A struct type implement Repository interface
type CustomStruct struct {
	Num int
}

func (c CustomStruct) Find(id int) error {
	return CustomError{message: strconv.Itoa(id)}
}

//A string type implement Repository interface
type CustomString string

func (c CustomString) Find(id int) error {
	return errors.New(strconv.Itoa(id))
}

//When return a Repository interface, Anything that implement this interface can be returned
//1. CustomStruct
func NewRepositoryStruct(db *sql.DB) Repository {
	return CustomStruct{9}
}

//2. CustomString
func NewRepositoryString(id int) Repository {
	a := CustomString(3)
	return a
}

//3. repository
func NewRepository(db *sql.DB) Repository {
	return repository{database: db}
}
