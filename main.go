// Package classification Go GIN API.
//
// This application demonstrates go-swagger
//
//     Schemes: http
//     Host: localhost:8080
//     BasePath: /test/v1
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(swaggerMiddleware())

	group := r.Group("/test/v1")

	group.GET("/ping", ping)
	group.GET("/users", getUsers)

	r.Run()

}

// swagger:model UserResponseParams
type UserResponseParams []UserResponseParam

// swagger:model UserResponseParam
type UserResponseParam struct {
	ID    string `json:"user_id"`
	Email string `json:"email"`
}

// swagger:operation GET /ping users get-ping
//
// Get ping
//
// Get ping
//
// ---
// consumes:
// - application/json
// produces:
// - application/json
// responses:
//   '200':
//     description: Successful operation

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// swagger:operation GET /users users get-users
//
// Get a list of users
//
// Get a list of users
//
// ---
// consumes:
// - application/json
// produces:
// - application/json
// parameters:
// - name: "MERCHANT-INFO-HEADER"
//   in: header
//   type: "string"
//   required: false
// - name: "access-token"
//   in: header
//   type: "string"
//   required: false
// - name: "refresh-token"
//   in: header
//   type: "string"
//   required: false
// responses:
//   '200':
//     description: Successful operation
//     schema:
//       "$ref": "#/definitions/UserResponseParams"
//   '401':
//     description: Unauthorized
func getUsers(c *gin.Context) {
	var users = UserResponseParams{
		{
			ID:    "1",
			Email: "1@test.com",
		},
	}

	c.JSON(http.StatusOK, users)
}
