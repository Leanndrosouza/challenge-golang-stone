package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

type scenarioCreateAccountTest struct {
	json           []byte
	statusExpected int
}

func TestCreateAccount(t *testing.T) {
	clearTable()

	jsonStr := []byte(`{"name":"Arthur Santos", "cpf": "777.970.100-05", "secret": "123456"}`)
	req, _ := http.NewRequest(http.MethodPost, "/accounts", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)

	var responseMap map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &responseMap)

	if responseMap["name"] != "Arthur Santos" {
		t.Errorf("Expected account name to be 'Arthur Santos'. Got '%v'", responseMap["name"])
	}

	if responseMap["cpf"] != "77797010005" {
		t.Errorf("Expected account cpf to be '77797010005'. Got '%v'", responseMap["cpf"])
	}

	// the balance is compared to 0.0 because JSON unmarshaling converts numbers to
	// floats, when the target is a map[string]interface{}
	if responseMap["balance"] != 0.0 {
		t.Errorf("Expected account balance to be '0'. Got '%v'", responseMap["balance"])
	}
}

func TestCreateAccountScenariosWhenMustFail(t *testing.T) {
	scenarios := []scenarioCreateAccountTest{
		{
			json:           []byte(`{}`),
			statusExpected: http.StatusBadRequest,
		},
		{
			json:           []byte(`{"cpf": "777.970.100-05", "secret": "123456"}`),
			statusExpected: http.StatusBadRequest,
		},
		{
			json:           []byte(`{"name":"Arthur Santos", "secret": "123456"}`),
			statusExpected: http.StatusBadRequest,
		},
		{
			json:           []byte(`{"name":"Arthur Santos", "cpf": "777.970.100-05"}`),
			statusExpected: http.StatusBadRequest,
		},
		{
			json:           []byte(`{"name":" ", "cpf": "777.970.100-05", "secret": "123456"}`),
			statusExpected: http.StatusBadRequest,
		},
		{
			json:           []byte(`{"name":"Antonio Vieira ", "cpf": "777.666.444-33", "secret": "123456"}`),
			statusExpected: http.StatusBadRequest,
		},
		{
			json:           []byte(`{"name":"Rafael Almeida Souza", "cpf": "777.970.100-05", "secret": "123456"}`),
			statusExpected: http.StatusCreated,
		},
		{
			json:           []byte(`{"name":"Vinicius Rabelo", "cpf": "77797010005", "secret": "123456"}`),
			statusExpected: http.StatusCreated,
		},
	}

	for _, scenario := range scenarios {
		clearTable()
		req, _ := http.NewRequest(http.MethodPost, "/accounts", bytes.NewBuffer(scenario.json))
		req.Header.Set("Content-Type", "application/json")

		response := executeRequest(req)
		checkResponseCode(t, scenario.statusExpected, response.Code)
	}
}

type scenarioGetBalanceTest struct {
	balance         int
	statusExpected  int
	generateAccount bool
	useInvalidParam bool
}

func TestGetAccountBalance(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	scenarios := []scenarioGetBalanceTest{
		{
			statusExpected:  http.StatusNotFound,
			generateAccount: false,
		},
		{
			statusExpected:  http.StatusBadRequest,
			generateAccount: false,
			useInvalidParam: true,
		},
		{
			balance:         rand.Intn(100000),
			statusExpected:  http.StatusOK,
			generateAccount: true,
		},
	}

	for _, scenario := range scenarios {
		clearTable()

		var accountID uint64
		var err error

		if scenario.generateAccount {
			accountID, err = insertAccountWithBalance(scenario.balance)
			if err != nil {
				t.Fatal(err)
			}
		}

		var req *http.Request

		if scenario.useInvalidParam {
			req, _ = http.NewRequest(http.MethodGet, "/accounts/invalid_param/balance", nil)
		} else {
			req, _ = http.NewRequest(http.MethodGet, fmt.Sprintf("/accounts/%d/balance", accountID), nil)
		}

		response := executeRequest(req)
		checkResponseCode(t, scenario.statusExpected, response.Code)

		var responseMap map[string]interface{}
		json.Unmarshal(response.Body.Bytes(), &responseMap)

		if scenario.generateAccount && responseMap["balance"] != float64(scenario.balance) {
			t.Errorf("Expected account balance to be '%v'. Got '%v'", scenario.balance, responseMap["balance"])
		}
	}
}
