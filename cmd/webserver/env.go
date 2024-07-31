package webserver

import (
	"os"
	"strings"
)

var (
	serverPort      string   = getEnv("SERVER_PORT")
	databaseURI     string   = getEnv("DATABASE_URI")
	databaseName    string   = getEnv("DATABASE_NAME")
	tankCollection  string   = getEnv("TANK_COLLECTION")
	stateCollection string   = getEnv("STATE_COLLECTION")
	kafkaTopic      string   = getEnv("KAFKA_TOPIC")
	kafkaBrokers    []string = parseEnv("KAFKA_BROKERS")
)

var defaultEnvs = map[string]string{
	"SERVER_PORT":      "8080",
	"DATABASE_URI":     "<INVALID>",
	"DATABASE_NAME":    "archimedes",
	"TANK_COLLECTION":  "tanks",
	"STATE_COLLECTION": "tank_states",
	"KAFKA_TOPIC":      "default",
}

var defaultArrayEnvs = map[string][]string{
	"KAFKA_BROKERS": {"localhost:9092"},
}

func getEnv(env string) string {
	value := os.Getenv(env)

	if value == "" {
		return defaultEnvs[env]
	}
	return value
}

func parseEnv(env string) []string {
	value := os.Getenv(env)

	if value == "" {
		return defaultArrayEnvs[env]
	}
	return strings.Split(value, ";")
}
