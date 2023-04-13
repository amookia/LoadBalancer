package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
)

func main() {
	node1, err := url.Parse("http://127.0.0.1:5001")
	if err != nil {
		log.Println(err)
	}
	node2, err := url.Parse("http://127.0.0.1:5002")
	if err != nil {
		log.Println(err)
	}
	nodes := []*url.URL{node1, node2}

	proxy := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rand := rand.Intn(2)
		r.Host = nodes[rand].Host
		r.URL.Host = nodes[rand].Host
		r.URL.Scheme = nodes[rand].Scheme
		r.RequestURI = ""
		response, err := http.DefaultClient.Do(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		w.WriteHeader(response.StatusCode)
		io.Copy(w, response.Body)
	})

	http.ListenAndServe(":8081", proxy)

}
