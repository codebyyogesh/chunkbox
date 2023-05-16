/*-----------------------------------------------------------
 @Filename:         routes.go
 @Copyright Author: Yogesh K
 @Date:             08/03/2023
-------------------------------------------------------------*/
package main

import "net/http"

// Update the signature for the routes() method so that it returns a
// http.Handler instead of *http.ServeMux.
func (app *application) routes() http.Handler{

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
    mux.HandleFunc("/chunkbox/view", app.chunkView)
    mux.HandleFunc("/chunkbox/create", app.chunkCreate)

   // Pass the servemux as the 'next' parameter to the secureHeaders middleware.
   // Because secureHeaders is just a function, and the function returns a
   // http.Handler we don't need to do anything else.
   // Wrap the existing chain with the logRequest middleware.
   // Middleware flow below
   // logRequest ↔ secureHeaders ↔ servemux ↔ application handler
    return app.logRequest(secureHeaders(mux))
}
