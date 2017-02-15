package main

import (
	"log"
	"net/http"
)

func main() {

	_, err := http.Get("https://www.sqliwithwrongDSN.com/")
	if err != nil {
		log.Fatal(err)
	}
}
