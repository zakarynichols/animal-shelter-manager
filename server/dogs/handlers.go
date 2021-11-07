package dogs

import (
	"encoding/json"
	"net/http"
	"server/utils"
)

type DogHandler interface {
	getDogById(handler DogStore) http.HandlerFunc
}

type dogHandler struct {
	DogHandler
}

func NewDogHandler() *dogHandler {
	return &dogHandler{}
}

func (handler *dogHandler) getDogById(store DogStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Content-Type", "application/json")

		var err error

		type Id struct {
			Id int `json:"dog_id"`
		}

		var id Id

		err = json.NewDecoder(r.Body).Decode(&id)

		if err != nil {
			utils.AppHttpError(w, utils.AppJsonError{Message: err.Error()}, http.StatusInternalServerError)
			return
		}

		dog, err := store.dog(id.Id)

		if err != nil {
			utils.AppHttpError(w, utils.AppJsonError{Message: err.Error()}, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

		err = json.NewEncoder(w).Encode(dog)

		if err != nil {
			utils.AppHttpError(w, utils.AppJsonError{Message: err.Error()}, http.StatusInternalServerError)
			return
		}
	}
}
