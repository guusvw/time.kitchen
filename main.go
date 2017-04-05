package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	kitchen := time.Now().Format(time.Kitchen)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", kitchen)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
