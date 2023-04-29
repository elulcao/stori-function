package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/elulcao/stori-function/pkg/email"
	"github.com/elulcao/stori-function/pkg/processor"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	data, err := processor.GetData("txns.csv")
	if err != nil {
		fmt.Fprint(w, "unable to get data")
		log.Print(err)
	}

	err = email.SendEmail(data)
	if err != nil {
		fmt.Fprint(w, "email not sent")
		log.Print(err)
	}

	message := "This HTTP triggered function executed successfully.\n"
	status := r.URL.Query().Get("status")
	if strings.EqualFold(status, "true") {
		fmt.Fprint(w, message)
	}
}

func main() {
	listenAddr := ":8080"

	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}

	http.HandleFunc("/api/processor", helloHandler)
	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
