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

		state, err := service.getState(req.Context(), input)
		if err != nil {
			app.GetLogger().Warn("error creating user", err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		respBytes, err := json.Marshal(state)
		if err != nil {
			app.GetLogger().Warn("error while marshilling users")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.Write(respBytes)
	})
}
