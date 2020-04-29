package main

import (
	"githubRepository/App"
	"log"
)

func main() {
	log.Print("Starting Server...")

	app := App.App{}

	app.Initialize()
}
