package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"zanko-khaledi/dsa/algorithmes"
)

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%v] %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

type Client struct {
	Requests []time.Time
}

type User struct {
	Id   int
	Name string
}

func main() {

	mux := http.NewServeMux()

	ratelimiterConf := &algorithmes.RatelimterConfig{
		Limit:    3,
		Window:   time.Second,
		Duration: 1,
	}

	handler := logger(ratelimiterConf.Ratelimiter(mux))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			data := []User{
				{
					Id:   1,
					Name: "zanko",
				},
			}

			b, _ := json.Marshal(data)

			w.WriteHeader(200)
			w.Write(b)
		}
	})

	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
