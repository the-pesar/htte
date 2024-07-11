package main

import (
	"fmt"
	"flag"
	"github.com/the-pesar/htte/htte"
)

func main() {
	port := flag.Int("p", 8080, "port to serve")

	flag.Parse()

	app := htte.New(htte.Configs{
		Port:    *port,
		Address: ":",
	})

	app.Get("/", func(req htte.Request) string {
		fmt.Println("handler")
		return "string"
	})

	app.Serve()
}
