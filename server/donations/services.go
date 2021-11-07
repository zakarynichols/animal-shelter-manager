package donations

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"server/utils"
	"strconv"

	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/charge"
)

func DonateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type Donation struct {
			Amount int64 `json:"donation_amount"`
		}

		var d Donation

		var err error

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		err = json.NewDecoder(r.Body).Decode(&d)

		if err != nil {
			utils.AppHttpError(w, utils.AppJsonError{Message: err.Error()}, http.StatusInternalServerError)
			return
		}

		donationAmount := d.Amount

		if err != nil {
			utils.AppHttpError(w, utils.AppJsonError{Message: err.Error()}, http.StatusInternalServerError)
			return
		}

		if donationAmount >= 50001 {
			utils.AppHttpError(w, utils.AppJsonError{Message: "Maximum donation amount is $500 USD"}, http.StatusBadRequest)
			return
		}

		params := &stripe.ChargeParams{
			Amount:      stripe.Int64(donationAmount),
			Currency:    stripe.String(string(stripe.CurrencyUSD)),
			Description: stripe.String("Test Description"),
			Source:      &stripe.SourceParams{Token: stripe.String("tok_mastercard")},
			Shipping: &stripe.ShippingDetailsParams{
				Address: &stripe.AddressParams{
					City: stripe.String(string("Test City")),
				},
				Name:           stripe.String("Test Name"),
				TrackingNumber: stripe.String(string(strconv.Itoa(rand.Int()))),
			},
		}

		c, err := charge.New(params)

		if err != nil {
			utils.AppHttpError(w, utils.AppJsonError{Message: err.Error()}, http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name: "id",
			MaxAge: 60,
			Value: "Test value",
			Path: "/",
		})
		
		w.WriteHeader(http.StatusOK)

		err = json.NewEncoder(w).Encode(&c)

		if err != nil {
			utils.AppHttpError(w, utils.AppJsonError{Message: err.Error()}, http.StatusInternalServerError)
			return
		}
	}
}

func ReadCookieHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		 // Read cookie
		 cookie, err := r.Cookie("id")
		 if err != nil {
			 fmt.Printf("Cant find cookie :/\r\n")
			 return
		 }
 
		 fmt.Printf("%s=%s\r\n", cookie.Name, cookie.Value)
	}
}