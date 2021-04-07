package controllers

import (
	"challenge-golang-stone/src/database"
	"challenge-golang-stone/src/models"
	"challenge-golang-stone/src/repositories"
	"challenge-golang-stone/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetAccounts return a list of accounts
func GetAccounts(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewAccountRepository(db)
	accounts, err := repository.GetAll()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, accounts)
}

// GetAccountBalance return a account balance
func GetAccountBalance(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	accountID, err := strconv.ParseUint(params["account_id"], 10, 64)
	if err != nil {
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
	account, err := repository.SearchByID(accountID)

	if err != nil {
		if err.Error() == "Account not found" {
			responses.Error(w, http.StatusNotFound, err)
		} else {
			responses.Error(w, http.StatusInternalServerError, err)
		}
		return
	}

	responses.JSON(w, http.StatusOK, map[string]int{
		"balance": account.Balance,
	})
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

	if err := account.Prepare("create"); err != nil {
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
