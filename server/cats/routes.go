package cats

import "github.com/gorilla/mux"

func Routes(router *mux.Router, query CatQuerier, handler CatHandler) {
	router.HandleFunc("/cats", handler.GetCatById(query)).Methods("GET")
}