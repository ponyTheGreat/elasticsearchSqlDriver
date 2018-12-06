package ge

import (
	"database/sql/driver"
	"io"
	"net/http"
)

// Rows contain the result
type Rows struct {
	ColumnsContent []column
	RowsContent    [][]interface{}
	Cur            int
	Closed         bool
	fetchSize      int
	cursor         string
	url            string
}

// Columns returns the names of the columns. The number of
// columns of the result is inferred from the length of the
// slice. If a particular column name isn't known, an empty
// string should be returned for that entry.
func (r *Rows) Columns() []string {
	var result []string
	for _, col := range r.ColumnsContent {
		result = append(result, col.Name)
	}
	return result
}

// Close closes the rows iterator.
//https://www.elastic.co/guide/en/elasticsearch/reference/6.5/sql-rest.html
func (r *Rows) Close() error {
	return nil
}

// Next is called to populate the next row of data into
// the provided slice. The provided slice will be the same
// size as the Columns() are wide.
//
// Next should return io.EOF when there are no more rows.
//
// The dest should not be written to outside of Next. Care
// should be taken when closing Rows not to modify
// a buffer held in dest.
func (r *Rows) Next(dest []driver.Value) error {
	if r.Cur >= len(r.RowsContent) {
		var result *http.Response
		var err error
		if r.cursor != "" {
			result, err = sendHTTPRequestCursor(r.cursor, r.url)
			if err != nil {
				return err
			}
			resultBytes, err := readHTTPResponse(result)
			if err != nil {
				return err
			}
			out, err := jsonDecode(resultBytes)
			if err != nil {
				return err
			}
			r.cursor = out.Cursor
			if out.Cursor == "" {
				return io.EOF
			}
			r.RowsContent = out.Rows
			r.Cur = 0
			for i, element := range (r.RowsContent)[r.Cur] {
				typeName := (r.ColumnsContent)[i].Coltype
				value, err := convertToValue(element, typeName)
				if err != nil {
					return err
				}
				dest[i] = value
			}
			r.Cur++
			// no more rows
		} else {
			r.Cur = 0
			return io.EOF
		}
		//current page does not run out
	} else {
		for i, element := range (r.RowsContent)[r.Cur] {
			typeName := (r.ColumnsContent)[i].Coltype
			value, err := convertToValue(element, typeName)
			if err != nil {
				return err
			}
			dest[i] = value
		}
		r.Cur++
	}
	return nil
}
