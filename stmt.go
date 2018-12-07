package ge

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// Stmt prepare
type Stmt struct {
	//method can be
	Method   string
	SQLQuery string
	URL      string
	Closed   bool
}

// DefaultFetchSize used when no fetch size inputed

//Close closes the stmt
func (stmt *Stmt) Close() error {
	if stmt.Closed {
		return errors.New("the stmt closed already")
	}
	stmt.Closed = true
	return nil
}

// NumInput output the number of arguments
//haven't implemented yet
func (stmt *Stmt) NumInput() int {
	return 0
}

// Exec executes a query that doesn't return rows, such
// as an INSERT or UPDATE.
//
// Deprecated: Drivers should implement StmtExecContext instead (or additionally).
func (stmt *Stmt) Exec(args []driver.Value) (driver.Result, error) {
	return nil, errors.New("Not implemented yet")
}

// Query executes a query that may return rows, such as a
// SELECT.
//
// Deprecated: Drivers should implement StmtQueryContext instead (or additionally).
func (stmt *Stmt) Query(args []driver.Value) (driver.Rows, error) {
	type requestBody struct {
		Query     string `json:"query"`
		FetchSize int    `json:"fetch_size"`
	}

	temp := requestBody{
		Query:     stmt.SQLQuery,
		FetchSize: 5,
	}

	request, err := json.Marshal(temp)
	if err != nil {
		return nil, err
	}

	//fetchSizeQuery := fmt.Sprintf("\"fetch_size\": %d", DefaultFetchSize)
	//requestBody := fmt.Sprintf("{\n  \"query\":\"%s\",\n  %s\n}", stmt.SQLQuery, fetchSizeQuery)

	//fmt.Println(string(request))
	resp, err := sendHTTPRequestQuery(request, "POST", stmt.URL)

	if err != nil {
		return nil, err
	}
	respBytes, err := readHTTPResponse(resp)
	var c outType
	if err := json.Unmarshal(respBytes, &c); err != nil {
		return nil, err
	}
	//fmt.Println(c.Columns[2].Coltype)
	result := getRows(c.Columns, c.Rows, c.Cursor, stmt.URL)
	//fmt.Println(&result)
	return &result, nil
}
