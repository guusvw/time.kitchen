package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/theckman/go-ipdata"
)

func main() {
	ipd := ipdata.NewClient("")

	http.HandleFunc("/", kitchen(ipd))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func kitchen(ipd ipdata.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f := r.Header.Get("x-forwarded-for")
		loc := time.UTC

		ip, err := ipd.Lookup(f)
		if err != nil {
			log.Printf("lookup failed: %v\n", err)
		} else {
			loc = ip.TimeZone
		}

		kitchen := time.Now().In(loc).Format(time.Kitchen)

		log.Printf("%s is in timezone %s: %s\n", f, loc.String(), kitchen)
		fmt.Fprint(w, kitchen)
	}
}
