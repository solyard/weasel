package weasel

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"
)

type Alerts struct {
	Alerts            []Alert                `json:"alerts"`
	CommonAnnotations map[string]interface{} `json:"commonAnnotations"`
	CommonLabels      map[string]interface{} `json:"commonLabels"`
	ExternalURL       string                 `json:"externalURL"`
	GroupKey          int                    `json:"groupKey"`
	GroupLabels       map[string]interface{} `json:"groupLabels"`
	Receiver          string                 `json:"receiver"`
	Status            string                 `json:"status"`
	Version           int                    `json:"version"`
}

type Alert struct {
	Annotations  map[string]interface{} `json:"annotations"`
	EndsAt       string                 `json:"endsAt"`
	GeneratorURL string                 `json:"generatorURL"`
	Labels       map[string]interface{} `json:"labels"`
	StartsAt     string                 `json:"startsAt"`
}

type BotConfig struct {
	BotToken string `yaml:"botToken"`
}

func LoadConfig() string {

	var cfg *BotConfig
	yamlFile, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return cfg.BotToken
}

func LoadTemplate() *template.Template {
	funcMap := template.FuncMap{
		"ToUpper": strings.ToUpper,
	}

	content, err := ioutil.ReadFile("config/default.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.New("Alert").Funcs(funcMap).Parse(string(content))
	if err != nil {
		fmt.Printf("Error while creating template: %v", err)
	}

	return tmpl
}
