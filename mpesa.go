package mpesa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func authenticate() string {
	var AuthResp AuthResponse

	client := Client()

	req, err := http.NewRequest("GET", "https://sandbox.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.SetBasicAuth(APP_KEY, APP_SECRET)
	req.Header.Add("cache-control", "no-cache")
	response, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	GetResponseBody(response, &AuthResp)
	return AuthResp.Access_token

}

func STKPushSimulation(s *STKPushSimulationD) {

	var i interface{}

	client := Client()

	b, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", "https://sandbox.safaricom.co.ke/mpesa/stkpush/v1/processrequest", bytes.NewBuffer(b))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("content-type", "application/json")
	req.Header.Set("authorization", "Bearer "+authenticate())
	req.Header.Set("cache-control", "no-cache")

	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	GetResponseBody(response, &i)
	fmt.Println(i)
}

func GetResponseBody(h *http.Response, i interface{}) {

	bodyText, err := ioutil.ReadAll(h.Body)
	if err != nil {
		log.Fatal(err)
	}

	err1 := json.Unmarshal(bodyText, &i)
	if err1 != nil {
		log.Fatal(err1)
	}
}

func Client() http.Client {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	return *client
}

// func STKPushTransactionStatus("174379","MTc0Mzc5YmZiMjc5ZjlhYTliZGJjZjE1OGU5N2RkNzFhNDY3Y2QyZTBjODkzMDU5YjEwZjc4ZTZiNzJhZGExZWQyYzkxOTIwMTcwODI0MTU1MDU1","20170824155055","ws_CO_27102017101215530");
