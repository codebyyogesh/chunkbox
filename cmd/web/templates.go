package main
import(
    "github.com/cpucortexm/chunkbox/internal/models"
 )
// Define a templateData type to act as the holding structure for
// any dynamic data that we want to pass to our HTML templates.
type templateData struct {
    Chunk *models.Chunk
    Chunks []*models.Chunk // Chunks field for holding a slice of chunks
}
