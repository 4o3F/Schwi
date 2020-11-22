package database

import (
	"database/sql"
	"fmt"
	"github.com/CardinalDevLab/Schwi-Backend/utils"
	_ "github.com/mattn/go-sqlite3"
)

var (
	MainDatabase *sql.DB
	PostDatabase *sql.DB
	err      error
)

func InitDatabase() {
	MainDatabase, err = sql.Open("sqlite3", "./data/databse/maindatabase.db")
	utils.ErrorHandler(err)

	PostDatabase, err = sql.Open("sqlite3", "./data/database/postdatabase.db")
	utils.ErrorHandler(err)

	fmt.Println("Init Database")
}
