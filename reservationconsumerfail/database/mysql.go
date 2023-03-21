package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Connector struct {
	Ad *sql.DB
}

type Adapter interface {
	CreateTable() error
	InsertTable(name, email string) error
	Close()
}

func NewConnector() Adapter {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/reservation")
	if err != nil {
		panic(err)
	} else if err = db.Ping(); err != nil {
		panic(err)
	}

	return &Connector{Ad: db}
}

func (c Connector) CreateTable() error {
	_, err := c.Ad.Exec("CREATE TABLE IF NOT EXISTS reservations (id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY, name TEXT NOT NULL, email TEXT NOT NULL)")
	if err != nil {
		return err
	}

	return nil
}

func (c *Connector) InsertTable(name, email string) error {
	// Create
	_, err := c.Ad.Exec("INSERT INTO reservations (name, email) VALUES (?,?)", name, email)
	if err != nil {
		return err
	}

	return nil
}

func (c *Connector) Close() {
	c.Ad.Close()
}
