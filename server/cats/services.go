package cats

import (
	"encoding/json"
	"net/http"
	"server/utils"
)

type Cat struct {
	CatId int    `json:"cat_id"`
	Name  string `json:"name"`
}

func GetCatById(handler CatQuerier) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Content-Type", "application/json")

		var err error

		type Id struct {
			Id int `json:"cat_id"`
		}

		var id Id

		err = json.NewDecoder(r.Body).Decode(&id)

		if err != nil {
			utils.AppHttpError(w, utils.AppJsonError{Message: err.Error()}, http.StatusInternalServerError)
			return
		}

		cat, err := handler.Cat(id.Id)

		if err != nil {
			utils.AppHttpError(w, utils.AppJsonError{Message: err.Error()}, http.StatusInternalServerError)
			return
		}

		
		w.WriteHeader(http.StatusOK)

		err = json.NewEncoder(w).Encode(cat)

		if err != nil {
			utils.AppHttpError(w, utils.AppJsonError{Message: err.Error()}, http.StatusInternalServerError)
			return
		}
	}
}