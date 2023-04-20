package handler

import (
	"io"
	"net/http"
	"net/url"

	"github.com/amookia/loadbalancer/internal/balancer"
	"go.uber.org/zap"
)

type service struct {
	nodes  []*url.URL
	logger *zap.SugaredLogger
}

type Handler interface {
	BalancerHandler(http.ResponseWriter, *http.Request)
	ErrorHandler(http.ResponseWriter, *http.Request)
}

func NewHandler(nodes []*url.URL, logger *zap.SugaredLogger) Handler {
	return &service{
		nodes:  nodes,
		logger: logger,
	}
}

func (h service) BalancerHandler(w http.ResponseWriter, r *http.Request) {
	b := balancer.Balancer{Nodes: h.nodes, Total: len(h.nodes)}
	response, err := b.PickNode(r)
	if err != nil {
		// handling errors
		// passing request to ErrorHandler
		h.ErrorHandler(w, r)
		return
	}
	h.logger.Infof("request proxied to : %s from : %s", r.Host, r.RemoteAddr)
	w.WriteHeader(response.StatusCode)
	io.Copy(w, response.Body)
}

// Just handling 500 error
func (h service) ErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
	io.WriteString(w, "<b>Error : 500</b>")
}
