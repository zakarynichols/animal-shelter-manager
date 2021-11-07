package users

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"server/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserHandler interface {
	Login(store UserStore) http.HandlerFunc
	Register(store UserStore) http.HandlerFunc
}

type userHandler struct {
	UserHandler
}

func NewUserHandler() *userHandler {
	return &userHandler{}
}

func (handler *userHandler) Login(store UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Content-Type", "application/json")
		// var err error

	}
}

func (handler *userHandler) Register(store UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Content-Type", "application/json")

		var err error

		var preAuthUser PreAuthenticatedUser

		err = json.NewDecoder(r.Body).Decode(&preAuthUser)

		if err != nil {
			utils.AppHttpError(w, utils.AppJsonError{Message: err.Error()}, http.StatusInternalServerError)
			return
		}

		err = store.canCreateUser(preAuthUser.Username)

		if err != nil {
			utils.AppHttpError(w, utils.AppJsonError{Message: err.Error()}, http.StatusConflict)
			return
		}

		bytes, err := bcrypt.GenerateFromPassword([]byte(preAuthUser.Password), 14)

		if err != nil {
			utils.AppHttpError(w, utils.AppJsonError{Message: err.Error()}, http.StatusInternalServerError)
			return
		}

		newSessionId, err := sessionId()

		if err != nil {
			utils.AppHttpError(w, utils.AppJsonError{Message: err.Error()}, http.StatusInternalServerError)
			return
		}

		newUser, err := store.createUser(preAuthUser.Username, bytes, newSessionId)

		if err != nil {
			utils.AppHttpError(w, utils.AppJsonError{Message: err.Error()}, http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:   "id",
			MaxAge: 60,
			Value:  newSessionId,
			Path:   "/",
		})

		w.WriteHeader(http.StatusOK)

		err = json.NewEncoder(w).Encode(newUser)

		if err != nil {
			utils.AppHttpError(w, utils.AppJsonError{Message: err.Error()}, http.StatusInternalServerError)
			return
		}
	}
}

// TODO: Expand on this. Middleware somehow?
// func IsAuthenticated(handlerFunc http.HandlerFunc) http.HandlerFunc {
// 	// return handlerFunc()
// 		//  cookie, err := r.Cookie("id")
// 		//  if err != nil {
// 		// 	return nil, err
// 		//  }
// 		//  fmt.Printf("%s=%s\r\n", cookie.Name, cookie.Value)
// 		//  return cookie, nil
// }

// TODO: Create secure session id
func sessionId() (string, error) {
	b := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
