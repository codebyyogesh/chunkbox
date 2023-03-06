/*-----------------------------------------------------------
 @Filename:         handlers.go
 @Copyright Author: Yogesh K
 @Date:             28/02/2023
-------------------------------------------------------------*/

package main

import (
    "fmt"
    "strconv"
    "net/http"
    "html/template"
    "log"
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
    // Initialize a slice containing the paths to the two files. It's important
    // to note that the file containing our base template must be the *first*
    // file in the slice.
    files := []string{
        "./ui/html/base.html",
        "./ui/html/partials/nav.html",
        "./ui/html/pages/home.html",
    }
    // Use the template.ParseFiles() function to read the template file into a
    // template set. If there's an error, we log the detailed error message and use
    // the http.Error() function to send a generic 500 Internal Server Error
    // response to the user.
    ts, err := template.ParseFiles(files...)
    if err != nil {
        log.Print(err.Error())
        http.Error(w, "Internal Server Error", 500)
        return
    }

    // Use the ExecuteTemplate() method to write the content of the "base" 
    // template as the response body.
    err = ts.ExecuteTemplate(w, "base", nil)
    if err != nil {
        log.Print(err.Error())
        http.Error(w, "Internal Server Error", 500)
    }
}

func chunkboxView(w http.ResponseWriter, r *http.Request){
    // Extract the value of the id parameter from the query string and try to
    // convert it to an integer using the strconv.Atoi() function. If it can't
    // be converted to an integer, or the value is less than 1, we return a 404 page
    // not found response.
    id, err :=  strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1{
        http.NotFound(w, r)
        return
    }
    // Use the fmt.Fprintf() function to interpolate the id value with our response
    // and write it to the http.ResponseWriter.
    // fmt.Fprintf() takes an io.Writer as the first parameter, but
    // we are passing w which is object of type ResponseWriter.
    // This can be done because ResponseWriter object satisfies the 
    // interface as it has a w.Write() method.
    fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func chunkboxCreate(w http.ResponseWriter, r *http.Request){
    // Use r.Method to check whether the request is using POST or not.
    if r.Method != http.MethodPost {
        // Use the Header().Set() method to add an 'Allow: POST' header to the
        // response header map. The first parameter is the header name, and
        // the second parameter is the header value.
        w.Header().Set("Allow", http.MethodPost)
        // Use the http.Error() function to send a 405 status code and "Method Not
        // Allowed" string as the response body.
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }
    w.Write([]byte("Create a small chunk..."))
}
