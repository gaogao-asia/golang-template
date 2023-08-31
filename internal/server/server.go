package server

import (
	"github.com/gaogao-asia/golang-template/pkg/connection"
)

func Run() {
	// create connection with database
	connection, err := connection.GetConnection()
	if err != nil {
		panic(err)
	}

	newHTTPServer(connection)
}
