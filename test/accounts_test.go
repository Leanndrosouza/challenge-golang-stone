package test

import (
	"bytes"
	"challenge-golang-stone/src/config"
	"challenge-golang-stone/src/database"
	"challenge-golang-stone/src/router"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	config.Load()

	ensureTableExists()

	code := m.Run()

	clearTable()

	os.Exit(code)
}

func ensureTableExists() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if _, err := db.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Exec("DELETE FROM accounts")
	db.Exec("ALTER SEQUENCE accounts_id_seq RESTART WITH 1")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS accounts
(
    id int auto_increment primary key,
    name varchar(50) not null,
    cpf varchar(11) not null unique,
    secret varchar(100) not null,
    balance int not null,
    created_at timestamp default current_timestamp()
)`

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	router := router.Generate()
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestCreateAccount(t *testing.T) {
	clearTable()

	db, err := database.Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

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

type scenarioTest struct {
	json           []byte
	statusExpected int
}

func TestCreateAccountScenariosWhenMustFail(t *testing.T) {
	db, err := database.Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	scenarios := []scenarioTest{
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
