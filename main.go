package main

import (
	"flag"
	"io"
	"log"
	"net/http"
)

const (
	certificateHeaderName string = "Cert"
)

var (
	upstreamHost = flag.String("upstream", "httpbin.org", "name of my upstream serive")
)

func main() {
	flag.Parse()
	http.HandleFunc("/", checkHeader)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func checkHeader(w http.ResponseWriter, r *http.Request) {
	cert, ok := r.Header[certificateHeaderName]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No cert!"))
		return
	}

	if len(cert) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("too many certs!"))
		return
	}

	log.Printf(cert[0])

	upstreamUrl := r.URL
	upstreamUrl.Host = *upstreamHost
	upstreamUrl.Scheme = "http"
	log.Printf("Upstream URL: %s", upstreamUrl.String())

	req, err := http.NewRequest(r.Method, upstreamUrl.String(), r.Body)
	if err != nil {
		log.Printf("Error in creating http.Request: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	req.Header = r.Header
	req.Header.Del(certificateHeaderName)

	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Couldn't call upstream service: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for k, vs := range resp.Header {
		for _, v := range vs {
			w.Header().Set(k, v)
		}
	}

	io.Copy(w, resp.Body)
}

func redirectPolicyFunc(r *http.Request, rs []*http.Request) error {
	return http.ErrUseLastResponse
}
