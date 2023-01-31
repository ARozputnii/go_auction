package main

import (
	"go_auction/internal/app/server"
)

func main() {
	s := server.App{}

	s.Start()
}
