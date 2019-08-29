package database

import (
	"database/sql"
	"log"
)

type Repository interface {
	Find(id int) error
}

type repository struct {
	// the db field is of type `DB` from the `database` package
	// that provides the method to interact with our real database
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

// Find retrieves a single record from database
// based on id column
func (r repository) Find(id int) error {}

// NewHandler constructs and returns a useable request handler.

func NewHandler(db *sql.DB,
	requestDuration *metrics.Histogram,
	logger *log.Logger, ) *Handler {

	return &Handler{

		db: db,
		dur: requestDuration,
		logger: logger,
	}
}
//dependency injections, give dependency to component.
handler = NewHandler(db, requestDur, logger)

func NewHandler(db dbtype, requestdur RequestDurType, logger, LoggerType)  &Handler{
	return Handler{
		db: db,
		requestdur: requ,
		logger: logger
}
}
