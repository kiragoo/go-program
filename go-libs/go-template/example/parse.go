package example

import (
	"html/template"
	"os"
)

type TelegrafConfStruct struct {
	InstanceID string `json:"instanceId"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Url        string `json:"url"`
	Http       string `json:"http"`
}

func GenerateTelegrafConf(t *TelegrafConfStruct) {
	tmpl, _ := template.ParseFiles("telegraf.conf.tpl")
	err := tmpl.Execute(os.Stdout, &t)
	if err != nil {
		panic(err)
	}
}
