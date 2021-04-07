package controllers

import (
	"challenge-golang-stone/src/auth"
	"challenge-golang-stone/src/database"
	"challenge-golang-stone/src/models"
	"challenge-golang-stone/src/repositories"
	"challenge-golang-stone/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// GetTransfers returns all transfers of a authenticated user
func GetTransfers(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	originID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	repository := repositories.NewTransferRepository(db)
	transfers, err := repository.GetByOriginID(originID)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, transfers)
}

// Transfer make a tranfer between two accounts
func Transfer(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var transfer models.Transfer

	if err = json.Unmarshal(body, &transfer); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	transfer.AccountOriginID, err = auth.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if err := transfer.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewTransferRepository(db)

	err = repository.CanTransfer(transfer)

	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if err = repository.Transfer(transfer); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, nil)
}
