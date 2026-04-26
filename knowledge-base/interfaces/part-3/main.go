package main

import (
	"fmt"
	"log"
	"part-3/mysqldb"
)

type dbcontract interface {
	Close()
	InsertUser(userName string) error
	SelectSingleUser(userName string) (string, error)
}

type Application struct {
	db dbcontract
}

func (this Application) Run() {
	userName := "exampleUser"
	err := this.db.InsertUser(userName)
	if err != nil {
		x := fmt.Errorf("Failed to insert user: %v", err)
		log.Printf("Failed to insert user: %v", x)
	}
}

func NewApplication(db dbcontract) *Application {
	return &Application{
		db: db,
	}
}
func main() {
	db, err := mysqldb.New("user", "password", "host", "port", "dbname")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	app := NewApplication(db)
	defer app.db.Close()
	app.Run()

}
