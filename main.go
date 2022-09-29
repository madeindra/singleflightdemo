package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/sync/singleflight"
)

// ab -n 100 -c 100  -t 10 http://localhost:8080/singleflight
func main() {
	var requestGroup singleflight.Group
	// localhost:8080/normal
	http.HandleFunc("/normal", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		// call function externalCall()
		status, err := externalCall()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// print processing time and status
		log.Printf("/normal handler requst: processing time: %+v | status %q", time.Since(start), status)

		fmt.Fprintf(w, "response Status: %q", status)
	})

	http.HandleFunc("/singleflight", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		v, err, shared := requestGroup.Do("apply_sf_key", func() (interface{}, error) {
			// call function external
			return externalCall()
		})

		// Check the error, as before.
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		status := v.(string)

		// print processing time and status
		if shared {
			log.Printf("/single flight handler requst: processing time: %+v | status %q | shared: %t", time.Since(start), status, shared)
		}
		fmt.Fprintf(w, "response Status: %q", status)
	})
	log.Println("running server on port :8080")
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func externalCall() (string, error) {
	time.Sleep(300 * time.Millisecond)
	log.Println("call external")
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("response externalCall: %s", resp.Status)
	}
	log.Println("response externalCall", resp.Status)
	return resp.Status, err
}
