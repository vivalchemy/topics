package main

import (
	"fmt"
	"io"
	"net/http"
)

const (
	baseURL = "http://localhost:8080"
)

func main() {
	// Create an HTTP client and execute the request
	client := &http.Client{}

	res, err := http.Get(baseURL)
	if err != nil {
		fmt.Errorf("ClientError: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body) // stores the jwt token in body
	if err != nil {
		fmt.Errorf("ClientError: %v", err)
	}

	token := string(body)

	// Verify the token
	req, err := http.NewRequest("GET", baseURL+"/verify", nil)
	if err != nil {
		fmt.Errorf("ClientError: %v", err)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	verRes, err := client.Do(req)
	if err != nil {
		fmt.Errorf("ClientError: %v", err)
	}
	defer verRes.Body.Close()

	// verify the jwt token
	if verRes.StatusCode != 200 {
		fmt.Errorf("ClientError: Invalid token")
	}

	body, err = io.ReadAll(verRes.Body)
	if err != nil {
		fmt.Errorf("ClientError: %v", err)
	}

	fmt.Println("verRes", string(body))

}
