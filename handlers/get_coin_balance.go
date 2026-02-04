package handlers

import (
	"encoding/json"
	"net/http"

	"go-learn/api"
	"go-learn/tools"

	log "github.com/sirupsen/logrus"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.CoinBalanceParams{}
	var err error

	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		params.Username = r.URL.Query().Get("username")
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalServerErrorHandler(w, err)
		return
	}

	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoins(params.Username)
	if tokenDetails == nil {
		log.Error(err)
		api.InternalServerErrorHandler(w, err)
		return
	}

	var response = api.CoinBalanceResponse{
		Balance: float64((*tokenDetails).Coins),
		Code:    http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalServerErrorHandler(w, err)
		return
	}
}
