package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"restJwt/models"
	userrepository "restJwt/repository/user"
	"restJwt/utils"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// Controller struct
type Controller struct{}

// Signup handler
func (c Controller) Signup(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("signup function")
		var user models.User
		var error models.Error
		json.NewDecoder(r.Body).Decode(&user)
		//spew.Dump(user)

		if user.Email == "" {
			error.Message = "Email is missing!"
			utils.RespondWithError(w, http.StatusBadRequest, error)
			return
		}

		if user.Password == "" {
			error.Message = "Password is missing!"
			utils.RespondWithError(w, http.StatusBadRequest, error)
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			log.Fatal(err)
		}

		user.Password = string(hash)

		userRepo := userrepository.UserRepository{}
		user = userRepo.Signup(db, user)

		if err != nil {
			error.Message = "Server Error"
			utils.RespondWithError(w, http.StatusInternalServerError, error)
			return
		}

		user.Password = "---Hidden----"

		utils.ResponseJSON(w, user)
	}
}

// Login controller
func (c Controller) Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("login function")
		var user models.User
		var error models.Error
		var jwt models.JWT

		json.NewDecoder(r.Body).Decode(&user)

		if user.Email == "" {
			error.Message = "Email is missing!"
			utils.RespondWithError(w, http.StatusBadRequest, error)
			return
		}

		if user.Password == "" {
			error.Message = "Password is missing!"
			utils.RespondWithError(w, http.StatusBadRequest, error)
			return
		}

		password := user.Password
		userRepo := userrepository.UserRepository{}
		user, err := userRepo.Login(db, user)

		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "models.User not found"
				utils.RespondWithError(w, http.StatusUnauthorized, error)
				return
			}
			log.Fatal(err)
		}

		hashedPassword := user.Password

		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

		if err != nil {
			error.Message = "Invalid password"
			utils.RespondWithError(w, http.StatusUnauthorized, error)
			return
		}

		token, err := utils.GenerateToken(user)

		if err != nil {
			log.Fatal(err)
		}

		jwt.Token = token

		w.WriteHeader(http.StatusOK)
		utils.ResponseJSON(w, token)
	}
}

// TokenVerifyMiddleWare controller
func (c Controller) TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var errorObject models.Error
		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) == 2 {
			authToken := bearerToken[1]

			token, error := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}

				return []byte("secret"), nil
			})

			if error != nil {
				errorObject.Message = error.Error()
				utils.RespondWithError(w, http.StatusUnauthorized, errorObject)
				return
			}

			if token.Valid {
				next.ServeHTTP(w, r)
			} else {
				errorObject.Message = error.Error()
				utils.RespondWithError(w, http.StatusUnauthorized, errorObject)
				return
			}
		} else {
			errorObject.Message = "Invalid token."
			utils.RespondWithError(w, http.StatusUnauthorized, errorObject)
			return
		}
	})
}
