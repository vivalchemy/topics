package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

type Middleware func(http.Handler) http.Handler

const (
	addr = ":8080"
)

func AttachMiddleware(f http.HandlerFunc, middleware ...Middleware) http.Handler {
	handler := http.Handler(f)
	for _, m := range middleware {
		handler = m(handler)
	}
	return handler
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Println("Req:", r.Proto, r.Method, r.URL.Path, start.Format("2006-01-02 15:04:05")) // this format is easier to filter
		next.ServeHTTP(w, r)
		fmt.Println("Resp:", r.Proto, r.Method, r.URL.Path, time.Since(start))
	})
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "text/plain")
	fmt.Fprint(w, "Pong")
}

// type HttpError struct {
// 	Status  int    `json:"status"`
// 	Message string `json:"message"`
// 	Error   string `json:"error"`
// }
//
// func CreateError(status int, message string, err error) HttpError {
// 	return HttpError{
// 		Status:  status,
// 		Message: message,
// 		Error:   err.Error(),
// 	}
// }

// Instead of using the types hardcoded we can just use an empty interface to
// inherit the types
// CreateError creates a structured error response
func CreateError(code int, message string, err error) map[string]interface{} {
	return map[string]interface{}{
		"status":  code,
		"message": message,
		"error":   err.Error(),
	}
}

func validateBody(body *User) error {
	if strings.TrimSpace(body.Name) == "" {
		return errors.New("Name is required")
	}
	if body.Age <= 0 {
		return errors.New("Age must be greater than 0")
	}
	if strings.TrimSpace(body.Email) == "" || !strings.Contains(body.Email, "@") {
		return errors.New("Email is required")
	}
	return nil
}

func bodyParsing(w http.ResponseWriter, r *http.Request) {
	// return the body as a string
	body := User{}

	w.Header().Set("Content-Type", "application/json")

	// parse the json body
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(CreateError(http.StatusBadRequest, "Invalid JSON", err))
		return
	}
	defer r.Body.Close()

	// validate the json body
	if err := validateBody(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(CreateError(http.StatusBadRequest, "Invalid Body", err))
		return
	}

	w.WriteHeader(http.StatusOK)
	str := fmt.Sprint("The name is ", body.Name, ", age is ", body.Age, ", email is ", body.Email)
	json.NewEncoder(w).Encode(struct {
		User
		Message string `json:"message"`
	}{
		User: User{
			Name:  body.Name,
			Age:   body.Age,
			Email: body.Email,
		},
		Message: str,
	})

	return
}

func formParsing(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	err := r.ParseForm() // use the url encoded form
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Access form values
	name := r.FormValue("name")
	age := r.FormValue("age")
	email := r.FormValue("email")

	// add a validation layer here
	for k, v := range r.Form {
		fmt.Println(k, v)
	}

	// Respond with the parsed data
	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Age: %s\n", age)
	fmt.Fprintf(w, "Email: %s\n", email)
}

func fileParsing(w http.ResponseWriter, r *http.Request) {
	receivedFile, details, err := r.FormFile("image") // image is the name of the form field
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer receivedFile.Close()
	fmt.Println("Received file:", details.Filename)

	// create a directory to store the file
	err = os.MkdirAll("uploads", 0755)
	if err != nil {
		fmt.Println(err)
		return
	}

	// create an empty file to store the file
	saveFileAt, err := os.Create(fmt.Sprintf("uploads/%v", details.Filename))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer saveFileAt.Close()

	// copy the uploaded file to the destination file
	_, err = io.Copy(saveFileAt, receivedFile) // it will always rewrite the file
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// respond with the file name
	fmt.Fprintf(w, "File uploaded successfully: %v", details.Filename)
}

func headerParsing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	headers := map[string]string{}
	for k, v := range r.Header {
		headers[k] = v[0]
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(headers)
}

// TODO: implement multipart parsing on own
func multipartParsing(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20) // 32 MB

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	fmt.Println("File:", handler.Filename)

	w.WriteHeader(http.StatusOK)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", ping)
	mux.HandleFunc("POST /body", bodyParsing)
	mux.HandleFunc("POST /form", formParsing)
	mux.HandleFunc("POST /file", fileParsing)
	// implement the get function for the file
	mux.HandleFunc("GET /header", headerParsing)
	mux.HandleFunc("POST /multipart", multipartParsing)
	// implement the get function for the big file using streaming
	mux.Handle("/middleware", AttachMiddleware(http.HandlerFunc(ping), Logger))
	// handle the sending of data for file and header

	fmt.Println("Hosted on http://localhost", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal("The service stopped working")
	}
}
