package main

import (
	"log"

	"github.com/vysmv/task-manager-api/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}