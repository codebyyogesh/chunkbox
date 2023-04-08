package models
import (
    "database/sql"
    "time"
    "errors"
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
    // Write the SQL statement we want to execute.
    stmt := `INSERT INTO chunks (title, content, created, expires)
    VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
    // Use the Exec() method on the embedded connection pool to execute the
    // statement. The first parameter is the SQL statement, followed by the
    // title, content and expiry values for the placeholder parameters. This
    // method returns a sql.Result type, which contains some basic
    // information about what happened when the statement was executed.
    result, err := m.DB.Exec(stmt, title, content, expires)
    if err != nil {
        return 0, err
    }
    // Use the LastInsertId() method on the result to get the ID of our
    // newly inserted record in the snippets table.
    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }
    // The ID returned has the type int64, so we convert it to an int type
    // before returning.
    return int(id), nil
}

// This will return a specific snippet based on its id.
func (m *ChunkModel) Get(id int) (*Chunk, error) {
    stmt := `SELECT id, title, content, created, expires FROM chunks
    WHERE expires > UTC_TIMESTAMP() AND id = ?`

    // Use the QueryRow() method on the connection pool to execute our
    // SQL statement, passing in the untrusted id variable as the value for the
    // placeholder parameter. This returns a pointer to a sql.Row object which
    // holds the result from the database.
    row := m.DB.QueryRow(stmt, id)

    // initialize a pointer to a new chunk struct
    c := &Chunk{}
    // Use row.Scan() to copy the values from each field in sql.Row to the
    // corresponding field in the Snippet struct. Notice that the arguments
    // to row.Scan are *pointers* to the place you want to copy the data into,
    // and the number of arguments must be exactly the same as the number of
    // columns returned by your statement.
    err := row.Scan(&c.ID, &c.Title, &c.Content, &c.Created, &c.Expires)

    if err != nil {
        // If the query returns no rows, then row.Scan() will return a
        // sql.ErrNoRows error. We use the errors.Is() function check for that
        // error specifically, and return our own ErrNoRecord error
        if errors.Is(err, sql.ErrNoRows) {
            return nil, ErrNoRecord
        } else {
            return nil, err
        }
    }
    // return chunk object
    return c, nil
}

// This will return the 10 most recently created snippets.
// We use slice of pointers to Chunk
func (m *ChunkModel) Latest() ([]*Chunk, error) {
    return nil, nil
}
