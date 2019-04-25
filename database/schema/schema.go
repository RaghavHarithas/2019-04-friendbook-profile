package schema

import (
	"github.com/RaghavHarithas/2019-04-friendbook-profile/database/connection"
)

var schema = `CREATE TABLE IF NOT EXISTS profile (
    email VARCHAR(50) PRIMARY KEY,
	username text,
	firstname text,
	lastname text,
    gender text,
	date_of_birth text,
	mobile_number text,
	city text,
	visibility text,
	status text
);`

func Init() {
	connection.ConnectPool()
	defer connection.DisconnectPool()

	connection.Pool.MustExec(schema)
}
