package main

import (
	"fmt"
	"log"
	"os/exec"

	p "github.com/promethues-loader/core/prometheus"
)

type PrometheusLoaderDefention struct {
	threads  int
	period   int
	ns       string
	sa       string
	interval int
	server 	 string
	headers  headers
}

type headers struct {
	Accept `json:"application/json, text/plain, */*"`
	Authorization `json:""`
	Accept-Encoding `json:"gzip, deflate, br"`
	Connection `json:"keep-alive"`
	X-Grafana-Org-Id `json:"1"`
}

func (pld *PrometheusLoaderDefention) getPrometheusInfo() {
	token, err := exec.Command(fmt.Sprintf("oc sa get-token %s -n %s", pld.ns, pld.sa)).Output()
	if err != nil {
		log.Fatal(err)
		exit(1)
	}
	pld.headers.Authorization = fmt.Sprintf("Bearer %s", token)

	route, err := exec.Command(fmt.Sprintf("oc get route %s -n %s |grep prometheus |awk '{print $2}'", pld.sa, self.ns)).Output()
	if err != nil {
		log.Fatal(err)
		exit(1)
	}
	pld.server = route
}

func prometheusLoader() {
}

func main() {
	dashboards := p.DashboardLoader()
	fmt.Println(dashboards)

	loader := PrometheusLoaderDefention{threads: 2, period: 2, ns: "promethues", sa: "", interval: 2}
	// start loader
}
