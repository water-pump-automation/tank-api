package webserver

import "os"

var (
	serverPort      = getEnv("SERVER_PORT")
	databaseURI     = getEnv("DATABASE_URI")
	databaseName    = getEnv("DATABASE_NAME")
	tankCollection  = getEnv("TANK_COLLECTION")
	stateCollection = getEnv("STATE_COLLECTION")
)

var defaultEnvs = map[string]string{
	"SERVER_PORT":         "8080",
	"DATABASE_URI":        "<INVALID>",
	"DATABASE_NAME":       "archimedes",
	"DATABASE_COLLECTION": "tanks",
}

func getEnv(env string) string {
	value := os.Getenv(env)

	if value == "" {
		return defaultEnvs[env]
	}
	return value
}
