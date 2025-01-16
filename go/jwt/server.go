package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	secret  = "my-secret"
	gIssuer = "creator"
)

var claims = jwt.MapClaims{
	// time needs to be resetted for each request so we need to add it at request
	// level
	"iss":      "creator",
	"username": "testuser",
}

// utils
func extractAuthorizationBearerToken(r *http.Request) (string, error) {
	token := r.Header.Get("Authorization")

	if token == "" {
		return "", errors.New("Authorization header not found")
	}
	if !strings.HasPrefix(token, "Bearer ") {
		return "", errors.New("Authorization header must start with Bearer")
	}

	token = strings.TrimPrefix(token, "Bearer ")
	return token, nil
}

func encodeJWT(claims jwt.MapClaims, secret string) (string, error) {
	claims["exp"] = time.Now().Add(time.Second * 120).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func decodeJWT(tokenString string, secret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// verify the algo
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		// verify the expiration time
		// if exp, ok := token.Claims.(jwt.MapClaims)["exp"].(float64); ok {
		// 	if time.Now().Unix() > int64(exp) {
		// 		return nil, errors.New("Token has expired")
		// 	}
		// } else {
		// 	return nil, errors.New("Expiration claim missing or invalid")
		// }
		if expirationTime, err := token.Claims.GetExpirationTime(); err != nil || expirationTime.Before(time.Now()) {
			return nil, errors.New("Token has expired")
		}

		// verify the issuer
		// if iss, ok := token.Claims.(jwt.MapClaims)["iss"].(string); !ok || iss != "creator" {
		// 	return nil, errors.New("Invalid issuer")
		// }
		if issuer, err := token.Claims.GetIssuer(); err != nil || issuer != gIssuer {
			return nil, errors.New("Invalid issuer")
		}

		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Could not parse claims")
	}

	return claims, nil
}

// controllers
func sendJWT(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/plain")

	token, err := encodeJWT(claims, secret)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
	}

	w.WriteHeader(http.StatusOK)

	// converting to url encoded format
	encodedToken := url.QueryEscape(token)
	fmt.Fprintf(w, encodedToken)
}

func verifyJWT(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/plain")

	// tokenString := r.URL.Query().Get("token")
	tokenString, err := extractAuthorizationBearerToken(r)
	// fmt.Println(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Error Encountered: %v", err)
		return
	}

	claims, err := decodeJWT(tokenString, secret)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Error Encountered: %v", err)
		return
	}

	// fmt.Fprintf(w, "Token is valid. %v", token.Claims.(jwt.MapClaims)["username"])
	if username, ok := claims["username"]; ok {
		fmt.Fprintf(w, "Token is valid. %v", username)
	} else {
		fmt.Fprintf(w, "Token is valid. No username given")
	}
}

// main
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", sendJWT)
	mux.HandleFunc("GET /verify", verifyJWT)

	fmt.Println("Server started on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Errorf("Error starting server: %s", err)
		panic(err)
	}
}
