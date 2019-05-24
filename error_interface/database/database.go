package database

import (
	"database/sql"
)

//Error
type CustomError struct {
	message string
}

func (e CustomError) Error() string {
	return e.message
}

func myMethod() error {
	return CustomError{message: "nono"}
}

type Repository interface {
	Find(id int) error
}

type CustomStruct struct {
	Num int
}

func (c CustomStruct) Find(id int) error {
	return CustomError{message: "dd"}
}

type repository struct {
	db *sql.DB
}
