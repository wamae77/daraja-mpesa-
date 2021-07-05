package mpesa

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func authenticateRequest(c *Credetials) string {
	var AuthResp AuthResponse

	client := Client()

	req, err := http.NewRequest("GET", endpoint+"/oauth/v1/generate?grant_type=client_credentials", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.SetBasicAuth(c.APP_KEY, c.APP_SECRET)
	req.Header.Add("cache-control", "no-cache")
	response, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	GetResponseBody(response, &AuthResp)
	return AuthResp.Access_token

}

func MpesaSTKPush(s *MpesaSTKPushBones, c *Credetials) (interface{}, error) {

	var i interface{}

	client := Client()

	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", endpoint+"/mpesa/stkpush/v1/processrequest", bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/json")
	req.Header.Set("authorization", "Bearer "+authenticateRequest(c))
	req.Header.Set("cache-control", "no-cache")

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	GetResponseBody(response, &i)
	return i, nil
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

func GenerateEncodedPassword(shortcode, passkey string) map[string]string {
	mytime := time.Now()
	formattedTime := mytime.Format("20060102150405")
	msg := shortcode + passkey + formattedTime
	encodedPassword := base64.StdEncoding.EncodeToString([]byte(msg))

	return map[string]string{
		"encoded":       encodedPassword,
		"formattedTime": formattedTime,
	}
}

func Client() http.Client {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	return *client
}
