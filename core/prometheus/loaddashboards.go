package prometheus

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

type Template struct {
	Name  string
	Query string
	Value string
}

type Dashboard struct {
	Name          string
	TargetQueries []string
	Annoation     string
	Templates     []Template
	Period        string
	Interval      string
}

type dashboardsDefenition struct {
	Items []struct {
		Data map[string]string `yaml:"data"`
	} `yaml:"items"`
}

type jsonDahboard struct {
	Rows []struct {
		Panels []struct {
			Targets []struct {
				Expr string `json:"expr"`
			} `json: "targets"`
		} `json: "panels"`
	} `json: "rows"`
	Templating struct {
		List []struct {
			Current struct {
				Value string `json: "value"`
			} `json: "current"`
			Name  string `json: "name"`
			Query string `json: "query"`
		} `json: "list"`
	} `json: "templating"`
	Time struct {
		From string `json: "from"`
	} `json: "time"`
	Refresh string `json: "refresh"`
}

func (yml *dashboardsDefenition) getyaml(yamlfile string) *dashboardsDefenition {
	var client http.Client
	resp, err := client.Get(yamlfile)
	if err != nil {
		// err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		err = yaml.Unmarshal(bodyBytes, yml)
		if err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}
	}
	return yml
}

func (jd *jsonDahboard) GetRefreshInterval() string {
	return jd.Refresh
}

func (jd *jsonDahboard) GetTimePeriod() string {
	return jd.Time.From
}

func (jd *jsonDahboard) GetTemplates() []Template {
	templates := make([]Template, len(jd.Templating.List))
	for _, template := range jd.Templating.List {
		tmp := Template{Name: template.Name, Query: template.Query, Value: template.Current.Value}
		templates = append(templates, tmp)
	}
	return templates
}

func (jd *jsonDahboard) GetTargets() []string {
	targets := make([]string, 0)
	for _, row := range jd.Rows {
		for _, panel := range row.Panels {
			for _, target := range panel.Targets {
				targets = append(targets, target.Expr)
			}
		}
	}
	return targets
}

func DashboardLoader() []Dashboard {
	var yaml string
	flag.StringVar(&yaml, "yaml", "https://raw.githubusercontent.com/openshift/cluster-monitoring-operator/master/assets/grafana/dashboard-definitions.yaml",
		"link to a dashboard definitation file")
	flag.Parse()

	var yml dashboardsDefenition
	yml.getyaml(yaml)

	dashboards := make([]Dashboard, 0, len(yml.Items))
	var jsondahboard jsonDahboard
	for _, dashboard := range yml.Items {
		for name, jsondata := range dashboard.Data {
			json.Unmarshal([]byte(jsondata), &jsondahboard)
			dashboards = append(dashboards, Dashboard{
				Name:          name,
				TargetQueries: jsondahboard.GetTargets(),
				Templates:     jsondahboard.GetTemplates(),
				Period:        jsondahboard.GetTimePeriod(),
				Interval:      jsondahboard.GetRefreshInterval(),
				Annoation:     "",
			})
		}
	}
	return dashboards
}
