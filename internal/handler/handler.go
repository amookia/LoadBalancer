package handler

import (
	"io"
	"math/rand"
	"net/http"
	"net/url"

	"go.uber.org/zap"
)

type service struct {
	nodes  []*url.URL
	logger *zap.SugaredLogger
}

type Handler interface {
	BalancerHandler(http.ResponseWriter, *http.Request)
}

func NewHandler(nodes []*url.URL, logger *zap.SugaredLogger) Handler {
	return &service{
		nodes:  nodes,
		logger: logger,
	}
}

func (h service) BalancerHandler(w http.ResponseWriter, r *http.Request) {
	rand := rand.Intn(2)
	r.Host = h.nodes[rand].Host
	r.URL.Host = h.nodes[rand].Host
	r.URL.Scheme = h.nodes[rand].Scheme
	h.logger.Infof("request proxied to : %s from : %s", r.Host, r.RemoteAddr)
	r.RequestURI = ""
	response, err := http.DefaultClient.Do(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.logger.Errorln("BalanceHandler", err.Error())
		return
	}
	w.WriteHeader(response.StatusCode)
	io.Copy(w, response.Body)
}
