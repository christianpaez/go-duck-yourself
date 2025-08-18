package main

import (
	"fmt"
	"net/http"
	"time"
)

const port = ":8080"

func logRequest(request *http.Request) {
	fmt.Printf("%s - %s %s\n", time.Now().Format("15:04:05"), request.Method, request.URL.Path)
}

func healthHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "ok\n")
	logRequest(request)
}

func main() {
	fmt.Printf("Server started at http://localhost%s\n", port)

	http.HandleFunc("/health", healthHandler)
	http.ListenAndServe(port, nil)
}
