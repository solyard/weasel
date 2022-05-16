package weasel

import (
	"io/ioutil"
	"strings"
	"text/template"

	log "github.com/sirupsen/logrus"

	"github.com/ilyakaznacheev/cleanenv"
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
	BotToken string `env:"TELEGRAM_BOT_TOKEN"`
}

func LoadConfig() string {
	var cfg BotConfig
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		log.Fatalf("Error while reading bot token from TELEGRAM_BOT_TOKEN variable. Error: %v", err)
	}
	if cfg.BotToken == "" {
		log.Fatal("Please set up bot token to run this app: %v", err)
	}
	return cfg.BotToken
}

func LoadTemplate() *template.Template {
	funcMap := template.FuncMap{
		"ToUpper": strings.ToUpper,
	}

	content, err := ioutil.ReadFile("config/default.tmpl")
	if err != nil {
		log.Fatalf("Error while reading template file: %v", err)
	}

	tmpl, err := template.New("Alert").Funcs(funcMap).Parse(string(content))
	if err != nil {
		log.Fatalf("Error while creating template: %v", err)
	}

	return tmpl
}
