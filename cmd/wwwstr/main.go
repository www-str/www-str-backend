package main

import (
	"log"
	"wwwstr/internal/wwwstr"
)

func main() {
	err := wwwstr.Listen()
	if err != nil {
		log.Fatal(err)
	}
}
