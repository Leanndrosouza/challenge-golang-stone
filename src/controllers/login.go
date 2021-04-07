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

	var login models.Login

	if err = json.Unmarshal(body, &login); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := login.Prepare(); err != nil {
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
	accountSaved, err := repository.SearchByCPF(login.Cpf)

	if err != nil {
		if err.Error() == "Account not found" {
			responses.Error(w, http.StatusNotFound, err)
		} else {
			responses.Error(w, http.StatusInternalServerError, err)
		}
		return
	}

	if err = security.VerifyPassword(accountSaved.Secret, login.Secret); err != nil {
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
