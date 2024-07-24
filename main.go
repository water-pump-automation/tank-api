package main

import (
	"os"
	"water-tank-api/cmd/webserver"
)

var (
	api = os.Getenv("SERVER_API")
)

func main() {
	if api == "INTERNAL" {
		webserver.Internal()
	} else {
		webserver.External()
	}
}
