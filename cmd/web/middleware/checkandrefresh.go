package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/wertick01/dclib/internals/app/models"
)

var mySigningKey = []byte("johenews")

func CheckToken(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request, error) {
	token, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return w, nil, err
		}
		w.WriteHeader(http.StatusBadRequest)
		return w, nil, err
	}

	tknStr := token.Value
	claims := &models.Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return w, nil, err
		}
		w.WriteHeader(http.StatusBadRequest)
		return w, nil, err
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return w, nil, err
	}

	return w, r, nil
}
