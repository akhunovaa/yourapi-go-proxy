package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		keys, ok := r.URL.Query()["url"]
		key := keys[0]

		if !ok || len(keys[0]) < 1 {
			log.Println("Url Param 'url' is missing")
			return
		}
		log.Println("Url Param 'key' is: " + string(key))
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path+" KEY: "+key))
	})
	log.Println("Listening on localhost:7733")
	log.Fatal(http.ListenAndServe(":7733", nil))
}
