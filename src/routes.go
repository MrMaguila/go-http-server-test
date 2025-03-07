package main

import (
	"fmt"
    "net/http"
	"os"
	"io/ioutil"
	"strings"
)


func callEndpoint(w http.ResponseWriter, r *http.Request) {
	var url string;
	var orinal_url = r.URL.String();
	var params_i int = strings.Index(orinal_url, "?") + 1;

	// build the request url string
	if params_i > 0 {
		url = REDIRECT_URL + "?" + orinal_url[params_i:];
	} else {
		url = REDIRECT_URL;
	}

	fmt.Printf("Requesting: %v", url);
	
	resp, err := http.Get(url);

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)

		w.WriteHeader(http.StatusInternalServerError);
		w.Header().Set("Content-Type", "application/json");

		return;
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", REDIRECT_URL, err)
		
		w.WriteHeader(http.StatusInternalServerError);
		w.Header().Set("Content-Type", "application/json");
		
		return;
	}

	w.Header().Set("Access-Control-Allow-Origin", "*");
	fmt.Fprintf(w, string(b[:]));
}
