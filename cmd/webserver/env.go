package webserver

import "os"

var (
	serverPort         = os.Getenv("SERVER_PORT")
	databaseURI        = os.Getenv("DATABASE_URI")
	databaseName       = os.Getenv("DATABASE_NAME")
	databaseCollection = os.Getenv("DATABASE_COLLECTION")
)
