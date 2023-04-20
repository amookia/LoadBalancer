package handler

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
)

type service struct {
	nodes []*url.URL
}

type Handler interface {
	BalancerHandler(http.ResponseWriter, *http.Request)
}

func NewHandler(nodes []*url.URL) Handler {
	return &service{
		nodes: nodes,
	}
}

func (h service) BalancerHandler(w http.ResponseWriter, r *http.Request) {
	rand := rand.Intn(2)
	r.Host = h.nodes[rand].Host
	r.URL.Host = h.nodes[rand].Host
	r.URL.Scheme = h.nodes[rand].Scheme
	r.RequestURI = ""
	response, err := http.DefaultClient.Do(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.WriteHeader(response.StatusCode)
	io.Copy(w, response.Body)
}
