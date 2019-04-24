package connection

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	host     = "profiledb"
	port     = "3306"
	user     = "root"
	password = "profile"
	dbName   = "profile"
)

type DB struct {
	*sqlx.DB
}

//Handle it with care, you may end up with multiple open connection to DB
var Pool *DB

func ConnectPool() {
	dbConnection := GetConnectionString()
	var err error
	var db *sqlx.DB

	for i := 0; i < 30; i++ {
		db, err = sqlx.Connect("mysql", dbConnection)
		if err == nil {
			Pool = &DB{db}
			return
		}
		log.Println("Reconnecting to DB")
		log.Printf("Error: %s", err)
		time.Sleep(5 * time.Second)
	}

	panic(err)
}

func DisconnectPool() {
	Pool.Close()
}

func GetConnectionString() string {
	dbConnection := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user,
		password,
		host,
		port,
		dbName,
	)

	return dbConnection
}
