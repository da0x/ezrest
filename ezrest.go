//
//   Copyright 2020 Justin Gehr
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License
//   You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2
//
//   Unless required by applicable law or agreed to in writing,
//   distributed under the License is distributed on an "AS IS" BASIS
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied
//   See the License for the specific language governing permissions
//   limitations under the License
//

// ezrest implements a simple approach to http calls that will embedd certain default
// headers like content type to application/json.
package ezrest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"time"
)

// RequestHeaders is a map which contains the headers that will be applied to each rest command.
// Default value is = DefaultHeaders().
var RequestHeaders = DefaultHeaders()

// DefaultHeaders returns a map with the default headers. You can use SetDefaultHeaders()
// to provide custom values.
// To add new headers while preserving the default headers use something like this:
//
// headers := rest.DefaultHeaders()
// headers["new-header"] = "value"
// rest.RequestHeaders = headers
func DefaultHeaders() map[string]string {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Connection"] = "close"
	return headers
}

// Verbose will turn on / off logging debugging information.
var Verbose = false

// Get calls a url and parses the json output and unmarshals it into a struct. It returns
// the http response code (if one is available) and an optional error.
// rest.RequestHeaders will be used for the request.
func Get(url string, response interface{}) (int, error) {
	if Verbose {
		log.Println("ezrest.Get(", url, ")")
	}
	request, err := http.NewRequest("GET", url, bytes.NewReader([]byte{}))
	if err != nil {
		return 0, err
	}
	for key, value := range RequestHeaders {
		request.Header.Set(key, value)
	}
	var client = &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(request)
	if err != nil {
		if resp != nil {
			return resp.StatusCode, err
		}
		return 0, err
	}
	read, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, err
	}
	resp.Body.Close()
	if resp.StatusCode >= 300 {
		return resp.StatusCode, fmt.Errorf("ezrest.Get() %v:%v", resp.StatusCode, string(read))
	}
	err = json.Unmarshal(read, response)
	if err != nil {
		if Verbose {
			log.Println("Response:", string(read))
		}
		return resp.StatusCode, err
	}
	return resp.StatusCode, nil
}

// Post calls a url with the POST method and sends the provided body. It will
// return a response code (if one is available) and an optional error. This
// function will also unmarshal the reponse in its response struct.
// rest.RequestHeaders will be used for the request.
func Post(url string, body, response interface{}) (int, error) {
	if Verbose {
		log.Println("ezrest.Post(", url, ")")
	}
	requestBody, err := json.Marshal(body)
	if err != nil {
		return 0, err
	}
	request, err := http.NewRequest("POST", url, bytes.NewReader(requestBody))
	if err != nil {
		return 0, err
	}
	for key, value := range RequestHeaders {
		request.Header.Set(key, value)
	}
	var client = &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(request)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	read, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, err
	}
	if len(read) > 0 && response != nil {
		err = json.Unmarshal(read, response)
		if err != nil {
			return resp.StatusCode, fmt.Errorf("Unmarshal error: %v, Response Body: %v", err, string(read))
		}
	}
	return resp.StatusCode, nil
}

// PostAcceptOctetStream will attempt to post the request to the given url and will give back
// the response from the server as a string.
// rest.RequestHeaders will be used for the request.
func PostAcceptOctetStream(url string, body interface{}, response *string) (int, error) {
	requestBody, err := json.Marshal(body)
	if err != nil {
		return 0, err
	}
	request, err := http.NewRequest("POST", url, bytes.NewReader(requestBody))
	if err != nil {
		return 0, err
	}
	for key, value := range RequestHeaders {
		request.Header.Set(key, value)
	}
	var client = &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(request)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	read, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, err
	}
	*response = string(read)
	return resp.StatusCode, nil
}
