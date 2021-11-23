//
//   Copyright 2020 Justin Gehr, Daher Alfawares
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

// DefaultHeaders returns a map with the default headers. You can use SetDefaultHeaders()
// as an initial value and then provide custom values.
//
// headers := rest.DefaultHeaders()
// headers["new-header"] = "value"
// Get(..., headers, ...)
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
func Get(url string, headers map[string]string, response interface{}) (int, error) {
	if Verbose {
		log.Println("ezrest.Get(", url, ")")
	}
	request, err := http.NewRequest("GET", url, bytes.NewReader([]byte{}))
	if err != nil {
		return 0, err
	}
	for key, value := range headers {
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

// PostOrPut calls a url with the POST or PUT method as specified in the first parameter
// and sends the provided body. It will return a response code (if one is available)
// and an optional error. This function will also unmarshal the reponse in its response struct.
func PostOrPut(method, url string, headers map[string]string, body, response interface{}) (int, error) {
	if Verbose {
		log.Println("ezrest.Post(", url, ")")
	}
	requestBody, err := json.Marshal(body)
	if err != nil {
		return 0, err
	}
	request, err := http.NewRequest(method, url, bytes.NewReader(requestBody))
	if err != nil {
		return 0, err
	}
	for key, value := range headers {
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

// Post calls a url with the POST method and sends the provided body. It will
// return a response code (if one is available) and an optional error. This
// function will also unmarshal the reponse in its response struct.
func Post(url string, headers map[string]string, body, response interface{}) (int, error) {
	return PostOrPut("POST", url, headers, body, response)
}

// MustPost calls Post and will handle any errors internally. If the call to Post()
// fails, then MustPost will exit the program and display an error message.
// If it succeeds it will return the http status code back as an int.
// Check net/http for status codes.
func MustPost(url string, headers map[string]string, body, response interface{}) int {
    code, err := Post(url, headers, body, response)
    if err != nil {
        log.Fatalln("ezrest.MustPost(", url, headers, body, "):", err)                                                                         
    }
    return code
}

// Put calls a url with the PUT method and sends the provided body. It will
// return a response code (if one is available) and an optional error. This
// function will also unmarshal the reponse in its response struct.
func Put(url string, headers map[string]string, body, response interface{}) (int, error) {
	return PostOrPut("PUT", url, headers, body, response)
}

// MustPut calls Put and will handle any errors internally. If the call to Put()
// fails, then MustPut will exit the program and display an error message.
// If it succeeds it will return the http status code back as an int.
// Check net/http for status codes.
func MustPut(url string, headers map[string]string, body, response interface{}) int {
    code, err := Put(url, headers, body, response)
    if err != nil {
        log.Fatalln("ezrest.MustPut(", url, headers, body, "):", err)                                                                         
    }
    return code
}

// PostAcceptOctetStream will attempt to post the request to the given url and will give back
// the response from the server as a string.
func PostAcceptOctetStream(url string, headers map[string]string, body interface{}, response *string) (int, error) {
	requestBody, err := json.Marshal(body)
	if err != nil {
		return 0, err
	}
	request, err := http.NewRequest("POST", url, bytes.NewReader(requestBody))
	if err != nil {
		return 0, err
	}
	for key, value := range headers {
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
