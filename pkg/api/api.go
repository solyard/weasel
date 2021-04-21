package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"

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
		log.Fatalf("%v", err)
	}

	var alertBody *weasel.Alerts
	_ = json.Unmarshal(body, &alertBody)

	//fmt.Printf("%s\n", body)
	//fmt.Printf("%v", *&alertBody.Alerts[0].Annotations)
	var responce bytes.Buffer
	err = AlertTemplate.Execute(&responce, alertBody)
	if err != nil {
		fmt.Printf("Error while executing template: %v", err)
	}

	telegram.SendMessageToBot(fmt.Sprintf("%s", &responce), vars["chat_id"])
}

func InitialiseAPI() {
	// Initializing Router
	fmt.Println("Launching API and Loading configuration")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc(APIEndpoint+"alert/{chat_id}", recieveAlertJSON).Methods("POST")

	AlertTemplate = weasel.LoadTemplate()
	fmt.Println("Configuration Loaded! Starting API ...")
	log.Fatal(http.ListenAndServe(":8081", router))
}
