package controllers

import (
	"fmt"
	"net/http"

	"github.com/jpadhye/restJwt/utils"
)

// ProtectedEndpoint handler
func (c Controller) ProtectedEndpoint() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("protected endpoint")
		utils.ResponseJSON(w, "YES")
	}
}
