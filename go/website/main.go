package main

// package website

import (
	"fmt"
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! this is me")
}

func getHtml(url string) {

	fmt.Println("Reached here")
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	bytesData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Reached here 2")
	fmt.Println(string(bytesData))
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listening on port :8080")
	http.ListenAndServe(":8080", nil)

	// getHtml("http://localhost:8080")
}
