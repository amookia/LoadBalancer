package balancer

import (
	"fmt"
	"net/http"
)

type Nodes struct {
	Response     *http.Response
	ResponseTime float64
}

type Judge struct {
	Nodes []Nodes
}

func (j Judge) ChooseURL() *http.Response {
	fmt.Println(j)
	return j.Nodes[1].Response
}
