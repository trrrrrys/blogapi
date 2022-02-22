package main

import (
	"log"
	"os"
)

func main() {
	server, cleanup, err := InitializeServer()
	if err != nil {
		panic(err)
	}
	defer cleanup()

	log.Printf("oO0OoO0OoO0Oo server start localhost%s oO0OoO0OoO0Oo\n", server.Addr)
	log.Printf("os.Getenv(PROJECT_ID): %v\n", os.Getenv("PROJECT_ID"))
	log.Println(server.ListenAndServe())
}
