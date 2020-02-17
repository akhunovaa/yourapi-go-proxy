package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		keys, ok := r.URL.Query()["url"]
		url := keys[0]

		if !ok || len(keys[0]) < 1 {
			log.Println("Url Param 'url' is missing")
		}
		log.Println("Url Param 'key' is: " + string(url))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		requestedBody := processServer(url)
		fmt.Fprintf(w, html.EscapeString(requestedBody))
	})
	log.Println("Listening on localhost:7733")
	log.Fatal(http.ListenAndServe(":7733", nil))
}

func processServer(s string) string {
	fmt.Println(s, " : started processing")
	resp, err := http.Get(s)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	responseBody := string(body)
	log.Println(string(responseBody))
	return responseBody
}
