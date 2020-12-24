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
	"fmt"
	"testing"
)

type getResponse struct {
	Status string `json:"status"`
	Data   []data `json:"data"`
}

type data struct {
	Id             string `json:"id"`
	EmployeeSalary string `json:"employee_salary"`
}

// TestApiGET will test rest.Get function is working as expected
func TestApiGET(t *testing.T) {
	var response getResponse
	statusCode, err := Get("http://dummy.restapiexample.com/api/v1/employees", &response)
	if err != nil {
		t.Errorf("Got err %v while trying rest.Get():", err)
		return
	}
	if statusCode != 200 {
		t.Errorf("Expected status code 200, got %d instead", statusCode)
		return
	}
	if response.Status != "success" {
		t.Errorf("Expected response.Status to be success, got %v instead", response.Status)
	}
	if response.Data[0].Id != "1" {
		t.Errorf("Expected response.Data[0].Id to be 1, got %s instead", response.Data[0].Id)
	}
}

// TestApiGET will test rest.Get function is working as expected
func TestApiGET404(t *testing.T) {
	var response getResponse
	_, err := Get("http://dummy.restapiexample.com/api/v1/eoyees", &response)
	if err == nil {
		t.Errorf("We expect an error, got nothing")
		return
	}
	fmt.Println(err)
}

type postBody struct {
	Name   string `json:"name"`
	Salary string `json:"salary"`
	Age    string `json:"age"`
}

type postResponse struct {
	Status string `json:"status"`
	Data   struct {
		Name   string `json:"name"`
		Salary string `json:"salary"`
		Age    string `json:"age"`
		Id     int    `json:"id"`
	} `json:"data"`
}

// TestApiPOST tests a POST operation
func TestApiPOST(t *testing.T) {
	var body postBody
	body.Name = "Joe Plum"
	body.Salary = "123"
	body.Age = "25"

	var response postResponse
	statusCode, err := Post("http://dummy.restapiexample.com/api/v1/create", body, &response)
	if err != nil {
		t.Errorf("Got err %v while trying to use api post", err)
		return
	}
	if statusCode != 200 {
		t.Errorf("Got status code 200 while trying to use api post, got %v instead", statusCode)
		return
	}
	if response.Status != "success" {
		t.Errorf("expected response.Status to be success, got %s instead", response.Status)
	}
	if response.Data.Name != "Joe Plum" {
		t.Errorf("expected response.Data.Name to be Joe Plum, got %s instead", response.Data.Name)
	}
	if response.Data.Salary != "123" {
		t.Errorf("expected response.Data.Salary to be 123, got %s instead", response.Data.Salary)
	}
	if response.Data.Age != "25" {
		t.Errorf("expected response.Data.Age to be 25, got %s instead", response.Data.Age)
	}
	return
}
