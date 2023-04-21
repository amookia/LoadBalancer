package balancer

import (
	"math/rand"
	"net/http"
	"net/url"
)

type Balancer struct {
	Nodes []*url.URL
	Total int
}

// PickNode
// This functions selects a node randomly
// If one node is damaged, select another one
func (b Balancer) PickNode(r *http.Request) (*http.Response, error) {
	rand := rand.Intn(b.Total)
	r.Host = b.Nodes[rand].Host
	r.URL.Host = b.Nodes[rand].Host
	r.URL.Scheme = b.Nodes[rand].Scheme
	r.RequestURI = ""
	response, err := http.DefaultClient.Do(r)
	for err != nil {
		response, err = b.PickNode(r)
	}
	return response, err
}
