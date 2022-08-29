package postgres

import (
	"database/sql"
	"encoding/json"

	_ "github.com/lib/pq"
)

type PGservice struct {
	uri string
	DB  *sql.DB
}

type PGError struct {
	Severity string `json:"Severity"`
	Code     string `json:"Code"`
	Message  string `json:"Message"`
}

func NewPGService(uri string) *PGservice {
	return &PGservice{
		uri: uri,
	}
}

func (pg *PGservice) Open() error {
	db, err := sql.Open("postgres", pg.uri)
	if err != nil {
		return err
	}

	pg.DB = db

	return nil
}

func (pg *PGservice) Close() error {
	return pg.DB.Close()
}

// NOTE: https://www.postgresql.jp/document/8.4/html/errcodes-appendix.html
func SQLState(err error) (string, error) {
	tmp, err := json.Marshal(err)
	if err != nil {
		return "", err
	}

	var pe PGError
	if err := json.Unmarshal(tmp, &pe); err != nil {
		return "", err
	}

	return pe.Code, nil
}
