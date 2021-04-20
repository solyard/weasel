package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/weasel/pkg/telegram"
	"github.com/weasel/pkg/weasel"
)

const APIEndpoint = "/api/v1/"

func recieveAlertJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("%v", err)
	}

	var alertBody *weasel.Alerts
	_ = json.Unmarshal(body, &alertBody)

	fmt.Printf("%s\n", body)
	telegram.SendMessageToBot(fmt.Sprintf("%s", body), vars["chat_id"])
	//fmt.Printf("%v", *&alertBody.Alerts[0].Annotations)
}

func InitialiseAPI() {
	// Initializing Router
	fmt.Print("Launching API ...")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc(APIEndpoint+"alert/{chat_id}", recieveAlertJSON).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", router))
}
