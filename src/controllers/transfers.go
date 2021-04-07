package controllers

import (
	"challenge-golang-stone/src/auth"
	"challenge-golang-stone/src/database"
	"challenge-golang-stone/src/repositories"
	"challenge-golang-stone/src/responses"
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

}
