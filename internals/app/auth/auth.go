package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/wertick01/dclib/internals/app/db"
	"github.com/wertick01/dclib/internals/app/models"
	"github.com/wertick01/dclib/internals/app/processors"
	"golang.org/x/crypto/bcrypt"
)

type Authoriser struct {
	processor *processors.UsersProcessor
}

func (m *Authoriser) hashPassword(s string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(s), 10)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func NewAuthoriser(processor *processors.UsersProcessor) *Authoriser {
	authoriser := new(Authoriser)
	authoriser.processor = processor
	return authoriser
}

var mySigningKey = []byte("johenews")

func (m *Authoriser) Login(w http.ResponseWriter, r *http.Request) {
	var creds *models.UserAuth

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := m.processor.FindByPhone(creds.UserPhone)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !db.ComparePassword(user.Hash, creds.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(1 * time.Minute)

	claims := &models.Claims{
		User: *user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func (m *Authoriser) Refresh(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tknStr := c.Value
	claims := &models.Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func (m *Authoriser) TimeChecker(w http.ResponseWriter, r *http.Request) {
	tk, _ := r.Cookie("token")
	vl := tk.Value

	c := make(chan os.Signal, 1)
	signal.Notify(c)

	ticker := time.NewTicker(30 * time.Second)
	stop := make(chan bool)

	go func() {
		defer func() { stop <- true }()
		for {
			select {
			case <-ticker.C:
				fmt.Println("30 seconds is gone...")
				fmt.Println(vl)
				m.Refresh(w, r)
			case <-stop:
				fmt.Println("Закрытие горутины")
				return
			}
		}
	}()

	<-c
	ticker.Stop()

	stop <- true
	<-stop
	fmt.Println("Приложение остановлено")
}
