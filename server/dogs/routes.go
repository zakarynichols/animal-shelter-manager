package dogs

import "github.com/gorilla/mux"

func Routes(router *mux.Router, store DogStore, handler DogHandler) {
	router.HandleFunc("/dogs", handler.getDogById(store)).Methods("GET")
}
