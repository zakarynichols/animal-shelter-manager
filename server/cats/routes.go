package cats

import "github.com/gorilla/mux"

func Routes(router *mux.Router, store CatStore, handler CatHandler) {
	router.HandleFunc("/cats", handler.getCatById(store)).Methods("GET")
}
