package http

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type HttpRequestHelper struct {
	Method        string
	URL           string
	RequestHeader map[string]string
	RequestBody   io.Reader
}

func (h *HttpRequestHelper) HttpRequest() string {
	var bodyString string

	req, err := http.NewRequest(h.Method, h.URL, h.RequestBody)
	if err != nil {
		log.Fatalf("new request error: %v", err)
	}

	if h.RequestHeader != nil {
		for k, v := range h.RequestHeader {
			req.Header.Add(k, v)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalf("request client error: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		bodyString = string(bodyBytes)
	} else {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString = string(bodyBytes)

		log.Printf("Response Status Code: %d", resp.StatusCode)
		log.Printf("Response Body: %s", bodyString)
	}

	return bodyString
}
