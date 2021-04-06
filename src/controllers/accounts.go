package controllers

import (
	"challenge-golang-stone/src/database"
	"challenge-golang-stone/src/models"
	"challenge-golang-stone/src/repositories"
	"challenge-golang-stone/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// GetAccounts return a list of accounts
func GetAccounts(w http.ResponseWriter, r *http.Request) {

}

// GetAccountBalance return a account balance
func GetAccountBalance(w http.ResponseWriter, r *http.Request) {

}

// CreateAccount insert a new account on database
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var account models.Account

	if err = json.Unmarshal(body, &account); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := account.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewAccountRepository(db)
	account.ID, err = repository.Create(account)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, account)
}
