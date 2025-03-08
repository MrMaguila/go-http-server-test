package main

import (
    "log"
    "net/http"
    "fmt"
)


func main() {
    use_ssl := PATH_TO_SSL_CERT_FILE != "" && PATH_TO_SSL_KEY_FILE != "";

    log.Println("Starting server...");
    http.HandleFunc("/main", callEndpoint);

    var err error;

    // Listen for requests
    if use_ssl {
        fmt.Println("listening on https://", fmt.Sprintf("%s:%s", ADDRESS_TO_BIND, PORT));
        err = http.ListenAndServeTLS(fmt.Sprintf("%s:%s", ADDRESS_TO_BIND, PORT), PATH_TO_SSL_CERT_FILE, PATH_TO_SSL_KEY_FILE, nil);
    } else {
        fmt.Println("listening on: http://", fmt.Sprintf("%s:%s", ADDRESS_TO_BIND, PORT));
        err = http.ListenAndServe(fmt.Sprintf("%s:%s", ADDRESS_TO_BIND, PORT), nil);
    }

    if err != nil {
        log.Fatal(err)
    }
}
