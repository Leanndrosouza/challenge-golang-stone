package test

import (
	"challenge-golang-stone/src/config"
	"challenge-golang-stone/src/database"
	"challenge-golang-stone/src/router"
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

	if _, err := db.Exec(tableAccountCreationQuery); err != nil {
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

const tableAccountCreationQuery = `CREATE TABLE IF NOT EXISTS accounts
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

func insertAccountWithBalance(balance int) (uint64, error) {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	statement, err := db.Prepare(
		"insert into accounts (name, cpf, secret, balance) values (?,?,?,?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec("Test Name", "37474338041", "$2a$10$K9kY1GTbYZuyzK3eo/l3FuVXun4bzCe2qBH/juEKZiPKu85u/mTB6", balance)
	if err != nil {
		return 0, err
	}

	lastIDInserted, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastIDInserted), nil
}
