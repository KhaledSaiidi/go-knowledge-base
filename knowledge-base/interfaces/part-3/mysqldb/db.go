package mysqldb

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	dbDriver = "mysql"
)

type Mysql struct {
	db *sql.DB
}

func New(dbUser, dbPassword, dbHost, dbPort, dbName string) (*Mysql, error) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Mysql{db: db}, nil
}

func (this Mysql) Close() {
	err := this.db.Close()
	if err != nil {
		log.Printf("mysqldb close failure: %v", err)
	}
}

func (this Mysql) InsertUser(userName string) error {
	_, err := this.db.Exec("INSERT INTO users (username) VALUES (?)", userName)
	return err
}

func (this Mysql) SelectSingleUser(userName string) (string, error) {
	var user string
	err := this.db.QueryRow("SELECT username FROM users WHERE username = ?", userName).Scan(&user)
	return user, err
}
