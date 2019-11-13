package main

import "mulbitchorong-backend/server"

func main() {
	s, err := server.New()
	if err != nil {
		panic(err)
	}
	s.Run(":3000")
}
