package donations

import "github.com/gorilla/mux"

func Routes(router *mux.Router, handler DonationHandler) {
	router.HandleFunc("/donate", handler.donate()).Methods("POST")
}