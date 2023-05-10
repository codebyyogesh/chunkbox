/*-----------------------------------------------------------
 @Filename:         handlers.go
 @Copyright Author: Yogesh K
 @Date:             28/02/2023
-------------------------------------------------------------*/

package main

import (
	"errors"
	"fmt"
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

    chunks, err := app.chunks.Latest()

    if err != nil{
        app.serverError(w, err)
        return
    }

    // Call the newTemplateData() helper to get a templateData struct containing
    // the 'default' data (which for now is just the current year), and add the
    // snippets slice to it.

    data := app.newTemplateData(r)
    data.Chunks = chunks

    // Use the render helper.
    app.render(w, 
               http.StatusOK,
               "home.html",
               data,
    )
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
    
    data := app.newTemplateData(r)
    data.Chunk = chunk

    // Use the render helper.
    app.render(w, 
               http.StatusOK,
               "view.html",
               data,
    )
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
    title := "On Rigveda"
    content := "One should, perform karma with nonchalance \n without expecting the benefits \n because sooner or later one shall definitely gets the fruits.. \n- Rigveda"
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
