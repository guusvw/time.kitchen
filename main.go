package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/theckman/go-ipdata"
)

func main() {
	ipd := ipdata.NewClient("")

	var timezones sync.Map

	http.HandleFunc("/", kitchen(ipd, timezones))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func kitchen(ipd ipdata.Client, timezones sync.Map) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		loc := time.UTC

		f := r.Header.Get("x-forwarded-for")
		if f != "" {
			tz, ok := timezones.Load(f)
			if ok {
				loc = tz.(*time.Location)
				log.Printf("loaded %s from timezones with %s\n", f, loc.String())
			} else {
				ip, err := ipd.Lookup(f)
				if err != nil {
					log.Println("failed to lookup ip", f, err)
				} else {
					loc = ip.TimeZone
					timezones.Store(f, loc)
					log.Printf("stored %s in timezones with %s\n", f, loc.String())
				}
			}
		}

		kitchen := time.Now().In(loc).Format(time.Kitchen)

		log.Printf("%s is in timezone %s: %s\n", f, loc.String(), kitchen)
		fmt.Fprint(w, kitchen)
	}
}
