package models
import (
    "database/sql"
    "time"
)
// define a chunk struct for an individual chunk.
// This will get stored in sql
type Chunk struct {
    ID      int
    Title   string
    Content string
    Created time.Time
    Expires time.Time
}
// Define a ChunkModel type which wraps a sql.DB connection pool.
type ChunkModel struct {
    DB *sql.DB
}
// This will insert a new snippet into the database.
func (m *ChunkModel) Insert(title string, content string, expires int) (int, error) {
    return 0, nil
}
// This will return a specific snippet based on its id.
func (m *ChunkModel) Get(id int) (*Chunk, error) {
    return nil, nil
}

// This will return the 10 most recently created snippets.
// We use slice of pointers to Chunk
func (m *ChunkModel) Latest() ([]*Chunk, error) {
    return nil, nil
}
