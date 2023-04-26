package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type db struct {
	conn *sql.DB
}

func (m *db) Ping() error {
	return m.conn.Ping()
}

func (m *db) Config() sql.DBStats {
	m.conn.SetMaxIdleConns(200)
	m.conn.SetConnMaxLifetime(time.Hour * time.Duration(24))
	m.conn.SetConnMaxIdleTime(time.Minute * time.Duration(24))
	m.conn.SetMaxOpenConns(500)
	return m.conn.Stats()
}

func (m *db) Close() error {
	return m.conn.Close()
}

func (m *db) Get() *sql.DB {
	return m.conn
}

func New(connStr string) (*db, error) {
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	fmt.Println("Ping: ", conn.Ping())

	return &db{conn: conn}, nil
}
