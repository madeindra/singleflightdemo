package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/sync/singleflight"
)

var requestGroup singleflight.Group

/*
how to test:
ab -n 100 -c 100 localhost:8080/normal
*/
func main() {
	// localhost:8080/normal
	http.HandleFunc("/normal", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		// call function githubStatus()
		status, err := githubStatus()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// print processing time and status
		log.Printf("/github handler requst: processing time: %+v | status %q", time.Since(start), status)

		fmt.Fprintf(w, "GitHub Status: %q", status)
	})

	http.HandleFunc("/singleflight", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		v, err, shared := requestGroup.Do("github", func() (interface{}, error) {
			// call function githubStatus()
			return githubStatus()
		})

		// Check the error, as before.
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		status := v.(string)

		// print processing time and status
		log.Printf("/github handler requst: processing time: %+v | status %q | shared: %t", time.Since(start), status, shared)

		fmt.Fprintf(w, "GitHub Status: %q", status)
	})
	log.Println("running server on port :8080")
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func githubStatus() (string, error) {
	time.Sleep(1 * time.Second)
	log.Println("call githubStatus")
	resp, err := http.Get("https://api.github.com")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("github response: %s", resp.Status)
	}

	return resp.Status, err
}
