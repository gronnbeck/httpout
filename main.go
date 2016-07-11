package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		byt, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte("An error occured"))
			return
		}
		defer r.Body.Close()

		fmt.Println(string(byt))

		w.Write([]byte("OK"))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
