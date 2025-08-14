package main

import (
	"Rest-root/internal/api/middlewares"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

type user struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	City string `json:"city"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello Root Route")
	w.Write([]byte("Hello Root Route"))
	fmt.Println("Hello Root Route")
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	//teachers/{id}
	//teachers/?key=value&query=value2&sortby=email&sortorder=ASC
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET Method on Teachers Route"))
		fmt.Println("Hello GET Method on Teachers Route")
	case http.MethodPost:
		w.Write([]byte("Hello POST Method on Teachers Route"))
		fmt.Println("Hello POST Method on Teachers Route")
	case http.MethodPut:
		w.Write([]byte("Hello PUT Method on Teachers Route"))
		fmt.Println("Hello PUT Method on Teachers Route")
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH Method on Teachers Route"))
		fmt.Println("Hello PATCH Method on Teachers Route")
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE Method on Teachers Route"))
		fmt.Println("Hello DELETE Method on Teachers Route")
	}
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET Method on Students Route"))
		fmt.Println("Hello GET Method on Students Route")
	case http.MethodPost:
		w.Write([]byte("Hello POST Method on Students Route"))
		fmt.Println("Hello POST Method on Students Route")
	case http.MethodPut:
		w.Write([]byte("Hello PUT Method on Students Route"))
		fmt.Println("Hello PUT Method on Students Route")
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH Method on Students Route"))
		fmt.Println("Hello PATCH Method on Students Route")
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE Method on Students Route"))
		fmt.Println("Hello DELETE Method on Students Route")
	}
}

func execsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET Method on Execs Route"))
		fmt.Println("Hello GET Method on Execs Route")
	case http.MethodPost:
		w.Write([]byte("Hello POST Method on Execs Route"))
		fmt.Println("Hello POST Method on Execs Route")
	case http.MethodPut:
		w.Write([]byte("Hello PUT Method on Execs Route"))
		fmt.Println("Hello PUT Method on Execs Route")
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH Method on Execs Route"))
		fmt.Println("Hello PATCH Method on Execs Route")
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE Method on Execs Route"))
		fmt.Println("Hello DELETE Method on Execs Route")
	}
}

func main() {
	port := ":3000"

	cert := "cert.pem"
	key := "key.pem"

	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)

	mux.HandleFunc("/teachers/", teachersHandler)

	mux.HandleFunc("/students/", studentsHandler)

	mux.HandleFunc("/execs/", execsHandler)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	//create custom server
	server := &http.Server{
		Addr: port,
		// Handler: mux,
		// Handler:   middlewares.SecurityHeaders(mux),
		Handler:   middlewares.Cors(mux),
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on port: ", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting the server", err)
	}
}
