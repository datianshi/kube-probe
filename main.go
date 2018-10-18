package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Value struct {
	Health bool
}

func main() {
	health := true
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {
			data, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Println(err)
				return
			}
			var v Value
			err = json.Unmarshal(data, &v)
			if err != nil {
				fmt.Println(err)
				return
			}
			health = v.Health
			w.WriteHeader(200)
			w.Write(data)
		} else {
			if !health {
				w.WriteHeader(500)
				w.Write([]byte("error"))
			} else {
				w.WriteHeader(200)
				w.Write([]byte("ok\n"))
			}
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
