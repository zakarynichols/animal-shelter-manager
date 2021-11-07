package users

import "github.com/gorilla/mux"

func Routes(router *mux.Router, store UserStore, handler UserHandler) {
	router.HandleFunc("/users/register", handler.Register(store)).Methods("POST", "OPTIONS")
	router.HandleFunc("/users/login", handler.Login(store)).Methods("POST", "OPTIONS")
}
