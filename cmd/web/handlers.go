/*-----------------------------------------------------------
 @Filename:         handlers.go
 @Copyright Author: Yogesh K
 @Date:             28/02/2023
-------------------------------------------------------------*/

package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/cpucortexm/chunkbox/internal/models"
)

// Start using the applications custom logger instead of the
// Go's standard logger. Update handler functions so that they become
// methods against the application struct.

func (app *application) home(w http.ResponseWriter, r *http.Request){
    // Check if the current request URL path exactly matches "/". If it doesn't, use
    // the http.NotFound() function to send a 404 response to the client.
    // Importantly, we then return from the handler. If we don't return, the handler
    // would keep executing and also write the "Hello from Chunkbox" message.
    if r.URL.Path != "/" {
        app.notFound(w) // use the app.notFound helper
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
        app.serverError(w,err) // Use the serverError() helper.
        return
    }

    // Use the ExecuteTemplate() method to write the content of the "base" 
    // template as the response body.
    err = ts.ExecuteTemplate(w, "base", nil)
    if err != nil {
        app.serverError(w,err) // Use the serverError() helper.
    }
}

func (app *application)chunkView(w http.ResponseWriter, r *http.Request){
    // Extract the value of the id parameter from the query string and try to
    // convert it to an integer using the strconv.Atoi() function. If it can't
    // be converted to an integer, or the value is less than 1, we return a 404 page
    // not found response.
    id, err :=  strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1{
        app.notFound(w) // use the app.notFound helper
        return
    }
    // Use the ChunkModel object's Get method to retrieve the data for a
    // specific record based on its ID. If no matching record is found,
    // return a 404 Not Found response.
    chunk, err := app.chunks.Get(id)

    if err != nil{
        if errors.Is(err, models.ErrNoRecord){
            app.notFound(w)
        }else {
            app.serverError(w, err)
        }
        return
    }

      // Write the snippet data as a plain-text HTTP response body.
    fmt.Fprintf(w, "%+v", chunk)

}

func (app *application)chunkCreate(w http.ResponseWriter, r *http.Request){
    // Use r.Method to check whether the request is using POST or not.
    if r.Method != http.MethodPost {
        // Use the Header().Set() method to add an 'Allow: POST' header to the
        // response header map. The first parameter is the header name, and
        // the second parameter is the header value.
        w.Header().Set("Allow", http.MethodPost)
        app.clientError(w, http.StatusMethodNotAllowed) // Use the clientError() helper.
        return
    }
    // Create some variables holding dummy data. We'll remove these later on
    // during the build.
    title := "On BhagvadGita"
    content := "The soul who meditates on the Self is content to serve the Self and rests satisfied within the Self; \n there remains nothing more for him to accomplish. \n- Bhagavad Gita 3.17"
    expires := 7
    // Pass the data to the ChunkModel.Insert() method, receiving the
    // ID of the new record back.
    id, err := app.chunks.Insert(title, content, expires)
    if err != nil {
        app.serverError(w, err)
        return
    }
    // Redirect the user to the relevant page for the chunk.
    http.Redirect(w, r, fmt.Sprintf("/chunkbox/view?id=%d", id), http.StatusSeeOther)

}
