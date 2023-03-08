package main

import "net/http"
// The routes() method returns a servemux containing our application routes.

func (app *application) routes() *http.ServeMux{

    // Initialise new server mux and register a home function
    // as handler for the "/" URL pattern
    mux := http.NewServeMux()

    // Create a file server which serves files out of the "./ui/static" directory.
    // Note that the path given to the http.Dir function is relative to the project
    // directory root.
    fileServer := http.FileServer(http.Dir("./ui/static/"))

    // Use the mux.Handle() function to register the file server as the handler for
    // all URL paths that start with "/static/". For matching paths, we strip the
    // "/static" prefix before the request reaches the file server.
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    mux.HandleFunc("/", app.home)
    mux.HandleFunc("/chunkbox/view", app.chunkboxView)
    mux.HandleFunc("/chunkbox/create", app.chunkboxCreate)

    return mux
}