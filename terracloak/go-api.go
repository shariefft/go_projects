package main

import (
	"fmt"
	"log"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
)

var jwtMiddleware *jwtmiddleware.JWTMiddleware

func main() {
	// Initialize JWT middleware for token verification
	jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("my-signing-key"), nil // Use the actual Keycloak public key
		},
		SigningMethod: jwt.SigningMethodRS256,
	})

	http.Handle("/protected", jwtMiddleware.Handler(http.HandlerFunc(protectedEndpoint)))
	http.HandleFunc("/public", publicEndpoint)

	fmt.Println("Server starting at :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func publicEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from public endpoint!"))
}

func protectedEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from protected endpoint!"))
}
