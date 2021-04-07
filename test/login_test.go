package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

type scenarioLoginTest struct {
	json           []byte
	statusExpected int
}

func TestLoginAccount(t *testing.T) {
	clearTable()

	// This account has CPF = 37474338041 and password = 123456
	_, err := insertAccountWithBalance(0)
	if err != nil {
		t.Fatal(err)
	}

	scenarios := []scenarioLoginTest{
		{
			json:           []byte(`{"cpf": "374.743.380-41", "secret": "123456"}`),
			statusExpected: http.StatusOK,
		},
		{
			json:           []byte(`{"cpf": "37474338041", "secret": "123456"}`),
			statusExpected: http.StatusOK,
		},
		{
			json:           []byte(`{}`),
			statusExpected: http.StatusBadRequest,
		},
		{
			json:           []byte(`{"name":"Arthur Santos", "secret": "123456"}`),
			statusExpected: http.StatusBadRequest,
		},
		{
			json:           []byte(`{"cpf": "777.970.100-05", "secret": "123456"}`),
			statusExpected: http.StatusNotFound,
		},
		{
			json:           []byte(`{"cpf": "37474338041", "secret": "111111"}`),
			statusExpected: http.StatusUnauthorized,
		},
	}

	for _, scenario := range scenarios {
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(scenario.json))
		req.Header.Set("Content-Type", "application/json")

		response := executeRequest(req)
		checkResponseCode(t, scenario.statusExpected, response.Code)
		if scenario.statusExpected == http.StatusOK {
			var responseMap map[string]interface{}
			json.Unmarshal(response.Body.Bytes(), &responseMap)

			if responseMap["token"] == nil {
				t.Errorf("Expected token to be setted. Got nil")
			}
		}
	}
}
