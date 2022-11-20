package state

import (
	"encoding/json"
	"kindlee/api"
	"kindlee/app"
	"net/http"
)

func GetSuitableState(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var input StateRequest
		err := json.NewDecoder(req.Body).Decode(&input)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		err = service.getState(req.Context(), input)
		if err != nil {
			app.GetLogger().Warn("error creating user", err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.Header().Add("Content-Type", "application/json")
		api.Success(rw, http.StatusOK, api.Response{Message: "User added Successfully"})
	})
}
