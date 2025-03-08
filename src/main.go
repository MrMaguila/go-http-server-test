package main

import (
    "log"
    "net/http"
    "fmt"
)


func main() {
    log.Println("Starting server...");
    use_ssl := PATH_TO_SSL_CERT_FILE != "" && PATH_TO_SSL_KEY_FILE != "";

    // Registering our handler functions
    http.HandleFunc("/main", callEndpoint);

    if use_ssl {
        fmt.Println("listening on https://", fmt.Sprintf("%s:%s", ADDRESS_TO_BIND, PORT));
    } else {
        fmt.Println("listening on: http://", fmt.Sprintf("%s:%s", ADDRESS_TO_BIND, PORT));
    }

    var err error;

    // Listen for requests
    if PATH_TO_SSL_CERT_FILE != "" && PATH_TO_SSL_KEY_FILE != "" {
        err = http.ListenAndServeTLS(fmt.Sprintf("%s:%s", ADDRESS_TO_BIND, PORT), PATH_TO_SSL_CERT_FILE, PATH_TO_SSL_KEY_FILE, nil);
    } else {
        err = http.ListenAndServe(fmt.Sprintf("%s:%s", ADDRESS_TO_BIND, PORT), nil);
    }

    if err != nil {
        log.Fatal(err)
    }
}
