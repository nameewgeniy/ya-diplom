package rest

import "net/http"

func (h RESTControllersImpl) UserBalanceWithdraw(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}
