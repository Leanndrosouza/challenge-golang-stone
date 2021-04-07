package controllers

import (
	"challenge-golang-stone/src/auth"
	"challenge-golang-stone/src/database"
	"challenge-golang-stone/src/models"
	"challenge-golang-stone/src/repositories"
	"challenge-golang-stone/src/responses"
	"challenge-golang-stone/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Login try to login user with email and password and after returns a web token
func Login(w http.ResponseWriter, r *http.Request) {
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

	if err := account.Prepare("login"); err != nil {
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
	accountSaved, err := repository.SearchByCPF(account.Cpf)

	if err != nil {
		if err.Error() == "Account not found" {
			responses.Error(w, http.StatusNotFound, err)
		} else {
			responses.Error(w, http.StatusInternalServerError, err)
		}
		return
	}

	if err = security.VerifyPassword(accountSaved.Secret, account.Secret); err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	token, err := auth.CreateToken(accountSaved.ID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, map[string]string{"token": token})
}
