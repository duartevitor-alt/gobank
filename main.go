package main

import (
	"fmt"
	"log"
)

func main() {

	fmt.Println("Starting our program")
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.CreateAccountTable(); err != nil {
		log.Fatal(err)
	}
	server := NewAPIServer(":3000", store)
	server.Run()

}
