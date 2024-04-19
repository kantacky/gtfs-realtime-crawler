package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("URL is not found. Input URL as an argument.")
	}
	url := os.Args[1]
	checkRenewalInterval(url)
}
