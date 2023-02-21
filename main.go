/*
-----------------------------------------------------------

	@Filename:         main.go
	@Copyright Author: Yogesh K
	@Date:             21/02/2023

-------------------------------------------------------------
*/
package main

import (
    "log"
    "net/http"
)
func home(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("Hello from Chunkbox"))
}


// server mux stores a mapping between the URL patterns for your
// application and the corresponding handlers.

func main() {
    // Initialise new server mux and register a home function
    // as handler for the "/" URL pattern
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)

    // Use the http.ListenAndServe() function to start a new web server. We pass in
    // two parameters: the TCP network address to listen on (in this case ":3001")
    // and the servemux we just created. If http.ListenAndServe() returns an error
    // we use the log.Fatal() function to log the error message and exit. Note
    // that any error returned by http.ListenAndServe() is always non-nil.
    log.Print("Starting server on :3001")
    err := http.ListenAndServe(":3001", mux)
    log.Fatal(err)
}


