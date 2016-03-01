package main

import "github.com/mattdotmatt/moodicle/server"

func main() {
	server.Start(8181, "./server/data")
}
