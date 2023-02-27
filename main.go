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
    // Check if the current request URL path exactly matches "/". If it doesn't, use
    // the http.NotFound() function to send a 404 response to the client.
    // Importantly, we then return from the handler. If we don't return, the handler
    // would keep executing and also write the "Hello from Chunkbox" message.
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    w.Write([]byte("Hello from Chunkbox"))
}

func chunkboxView(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("Display a small chunk..."))
}

func chunkboxCreate(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("Create a small chunk..."))
}


// We dont use DefaultServeMux because it is a global variable, 
// any package can access it and register a route — including any third-party
// packages that your application imports. If one of those third-party 
// packages is compromised, they could use DefaultServeMux to expose 
// a malicious handler to the web.

// server mux stores a mapping between the URL patterns for your
// application and the corresponding handlers. The server mux created
// here is a local one, unlike the DefaultServeMux

func main() {
    // Initialise new server mux and register a home function
    // as handler for the "/" URL pattern
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)
    mux.HandleFunc("/chunkbox/view", chunkboxView)
    mux.HandleFunc("/chunkbox/create", chunkboxCreate)

    // Use the http.ListenAndServe() function to start a new web server. We pass in
    // two parameters: the TCP network address to listen on (in this case ":3001")
    // and the servemux we just created. If http.ListenAndServe() returns an error
    // we use the log.Fatal() function to log the error message and exit. Note
    // that any error returned by http.ListenAndServe() is always non-nil.
    log.Print("Starting server on :3001")
    err := http.ListenAndServe(":3001", mux)
    log.Fatal(err)
}


