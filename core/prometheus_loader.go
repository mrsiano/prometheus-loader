package main

import (
	"crypto/tls"
	"net/http"
	"fmt"
	"log"
	"os/exec"

	p "github.com/promethues-loader/core/prometheus"
)

type PrometheusLoaderDefention struct {
	interval int
	period   int
	stepping int
	threads  int
	ns       string
	sa       string
	server 	 string
	query 	 string
	token
}

func (pld *PrometheusLoaderDefention) getPrometheusInfo() {
	token, err := exec.Command(fmt.Sprintf("oc sa get-token %s -n %s", pld.ns, pld.sa)).Output()
	if err != nil {
		log.Fatal(err)
		exit(1)
	}
	pld.token = fmt.Sprintf("Bearer %s", token)

	route, err := exec.Command(fmt.Sprintf("oc get route %s -n %s |grep prometheus |awk '{print $2}'", pld.sa, self.ns)).Output()
	if err != nil {
		log.Fatal(err)
		exit(1)
	}
	pld.server = route
}

func (pld *PrometheusLoaderDefention) requestGenerator(q) string {
	defer Mutex.unlock
	time_from := time.Now().Add(time.Minute * -pld.period).Unix()
	time_now :=  time.Now().Unix()
	pld.query = q
	return fmt.Sprintf("https://%s/api/v1/query_range?query=%s&start=%s&end=%s&step=%s",
		pld.server, q, time_from, time_now, pld.stepping)
}

func (pld *PrometheusLoaderDefention) request(rqst) {
	reqLogInfo := fmt.Sprintf("[%s] - concurrency:%s - query:%s",
		pld.dashboard.Name,
		pld.threads,
		pld.query)

		req, err := http.NewRequest("GET", rqst, nil)
		if err != nil {
				log.Fatal(err)
		}

		req.Header = map[string][]string{
			"Accept": "application/json, text/plain, */*",
			"Authorization": {pld.token},
			"Accept-Encoding": {"gzip, deflate, br"},
			"Connection string": {"keep-alive"},
			"X-Grafana-Org-Id": {"1"},
		}

		// This transport is what's causing unclosed connections.
		tr := &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		hc := &http.Client{Timeout: 2 * time.Second, Transport: tr}

		resp, err := hc.Do(req)
		if err != nil {
				log.Fatal(err)
		}
		defer resp.Body.Close()
}

func prometheusLoader() {
}

func main() {
	dashboards := p.DashboardLoader()
	fmt.Println(dashboards)

	loader := PrometheusLoaderDefention{threads: 2, period: 2, ns: "promethues", sa: "", interval: 2}
	// start loader
}
