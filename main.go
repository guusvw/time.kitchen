package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", kitchen)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func kitchen(w http.ResponseWriter, r *http.Request) {
	ip := r.Header.Get("x-forwarded-for")
	fmt.Println(ip)

	kitchen := time.Now().Format(time.Kitchen)
	fmt.Fprint(w, kitchen)
}
