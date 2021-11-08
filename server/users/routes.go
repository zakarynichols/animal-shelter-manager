package users

import "github.com/gorilla/mux"

func Routes(router *mux.Router, store UserStore, handler UserHandler) {
	router.HandleFunc("/users/register", handler.register(store)).Methods("POST", "OPTIONS")
	router.HandleFunc("/users/login", handler.login(store)).Methods("POST", "OPTIONS")
}