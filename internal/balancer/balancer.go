package balancer

import (
	"errors"
	"math/rand"
	"net/http"
	"net/url"
)

type Balancer struct {
	Nodes   []*url.URL
	Checked []*http.Response
	Total   int
	Count   int
}

// PickNode
// This functions selects a node randomly
// If one node is damaged, select another one
func (b Balancer) PickNode(r *http.Request) (*http.Response, error) {
	b.CheckNodes(r)
	rand := rand.Intn(len(b.Checked))
	if len(b.Checked) == 0 {
		return nil, errors.New("there is no node to response")
	}
	// counting error connections
	return b.Checked[rand], nil
}
