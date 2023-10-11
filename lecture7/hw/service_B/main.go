package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Data from Service B")
	})

	http.HandleFunc("/getA", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://localhost:8080/data")
		if err != nil {
			http.Error(w, "Failed to get data from Service A", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Failed to read response body", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Data from Service A: %s", body)
	})

	http.ListenAndServe(":8081", nil)
}
