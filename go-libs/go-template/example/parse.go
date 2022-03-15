package example

import (
	"html/template"
	"os"
)

type TelegrafConf struct {
	MetricConfig `json:"metricConfig,omitempty"`
	LogType      string `json:"logType,omitempty"`
	LogConfig    `json:"logConfig,omitempty"`
}

type MetricConfig struct {
	BCMetricGateway string `json:"bcMetricGateway"`
}

type LogConfig struct {
	EsUrl       string `json:"esUrl"`
	EsUsername  string `json:"esUsername"`
	EsPassword  string `json:"esPassword"`
	EsIndexName string `json:"esIndexName"`
}

type Task struct {
	Namespace       string `json:"namespace,omitempty"`
	EmqxApiUsername string `json:"emqxApiUsername,omitempty"`
	EmqxApiPassword string `json:"emqxApiPassword"`
	TelegrafConf    `json:"telegrafConf,omitempty"`
}

func GenerateTelegrafConf(t Task) {
	tmpl := template.Must(template.ParseFiles("template/telegraf.conf.tpl"))
	err := tmpl.Execute(os.Stdout, t)
	if err != nil {
		panic(err)
	}
}
