package main

import (
	"os"
	"water-tank-api/cmd/webserver"
)

var (
	api = getEnv("SERVER_API")
)

var defaultEnvs = map[string]string{
	"SERVER_API": "EXTERNAL",
}

func getEnv(env string) string {
	value := os.Getenv(env)

	if value == "" {
		return defaultEnvs[env]
	}
	return value
}

func main() {
	if api == "INTERNAL" {
		webserver.Internal()
	} else {
		webserver.External()
	}
}
