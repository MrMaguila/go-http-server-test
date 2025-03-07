package main

import (
    "log"
    "net/http"
)

const PORT string = ":8080"

func main() {
    log.Println("Starting server...")

    // Registering our handler functions
    http.HandleFunc("/main", callEndpoint)

    log.Println("Started on port", PORT)

    var err error;

    // Listen for requests
    if PATH_TO_SSL_CERT_FILE != "" && PATH_TO_SSL_KEY_FILE != "" {
        err = http.ListenAndServeTLS(PORT, PATH_TO_SSL_CERT_FILE, PATH_TO_SSL_KEY_FILE, nil);
    } else {
        err = http.ListenAndServe(PORT, nil);
    }

    if err != nil {
        log.Fatal(err)
    }
}
