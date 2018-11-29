package ge

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

//Conn for manipulate DB
type conn struct {
	host    string
	portnum string
}

//Prepare to produce a Stmt
//need to implemt check method for url, and sql query
//so far the only Http method is POST
func (c *conn) Prepare(query string) (driver.Stmt, error) {
	stmt := &Stmt{
		Method:   "POST",
		SQLQuery: query,
		URL:      fmt.Sprintf(fmt.Sprintf("http://%s:%s", c.host, c.portnum)),
		Closed:   false,
	}
	return stmt, nil
}

//Close closes the Conn
func (c *conn) Close() error {

	return errors.New("Not implemented yet")
}

//Begin a Tx
func (c *conn) Begin() (driver.Tx, error) {
	return nil, errors.New("Not implemented yet")
}

// Query to query DB with result
func (c *conn) Query(query string) (*driver.Rows, error) {
	stmt, err := c.Prepare(query)
	if err != nil {
		panic(err)
	}
	result, err := stmt.Query(nil)
	if err != nil {
		panic(err)
	}
	return &result, nil
}
