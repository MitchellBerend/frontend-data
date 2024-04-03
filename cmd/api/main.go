package main

import (
	"fmt"
	"frontend-data/internal/server"
	"log"
	"os"
)

func main() {
	server := server.NewServer()

	log.Printf("Listening on 0.0.0.0:%s\n", os.Getenv("PORT"))
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
