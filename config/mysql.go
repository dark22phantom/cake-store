package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	host        string
	port        string
	username    string
	password    string
	dbName      string
	maxIdlePool int
	maxIdleTime int

	DB *sql.DB
}

func (m *MySQL) Initialize() {
	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		m.username,
		m.password,
		m.host,
		m.port,
		m.dbName,
	)

	db, err := sql.Open("mysql", uri)
	if err != nil {
		log.Fatal(err)
	}
	m.DB = db

	err = m.DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MySQL!")
}
