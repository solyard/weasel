package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/weasel/pkg/telegram"
	"github.com/weasel/pkg/weasel"
)

const APIEndpoint = "/api/v1/"

var AlertTemplate *template.Template

func recieveAlertJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Error while reading body: %v", err)
		return
	}

	var alertBody *weasel.Alerts
	err = json.Unmarshal(body, &alertBody)
	if err != nil {
		log.Errorf("Error while unmarshaling JSON: %v", err)
		return
	}

	var response bytes.Buffer
	err = AlertTemplate.Execute(&response, alertBody)
	if err != nil {
		log.Errorf("Error while executing template: %v", err)
		return
	}

	telegram.SendAlert(fmt.Sprintf("%s", &response), vars["chat_id"])
}

func Start() {
	fmt.Println("Launching API and Loading configuration")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc(APIEndpoint+"alert/{chat_id}", recieveAlertJSON).Methods("POST")

	AlertTemplate = weasel.LoadTemplate()
	log.Info("Configuration Loaded! Starting API ...")
	log.Fatal(http.ListenAndServe(":8081", router))
}
