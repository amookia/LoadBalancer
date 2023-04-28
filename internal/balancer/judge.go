package balancer

import (
	"net/http"
)

func (b *Balancer) CheckNodes(r *http.Request) {
	for _, node := range b.Nodes {
		r.Host = node.Host
		r.URL.Host = node.Host
		r.URL.Scheme = node.Scheme
		r.RequestURI = ""
		response, err := http.DefaultClient.Do(r)
		if err == nil {
			b.Checked = append(b.Checked, response)
		}
	}
}
