package main

import (
	"log"
	"mulbitchorong-backend/server"
)

func main() {
	s, err := server.New()
	if err != nil {
		log.Fatal(err)
	}
	s.Run(":3000")
}
