package elasticsearchSqlDriver

import (
	"database/sql"
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
func (c *conn) Prepare(query string) (driver.Stmt, error) {
	//so far the only Http method is POST
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
func (c *conn) Query(query string) (*sql.Rows, error) {
	stmt, err := c.Prepare(query)
	if err != nil {
		panic(err)
	}
	result, err := stmt.Query(nil)
	if err != nil {
		panic(err)
	}

	re := &sql.Rows{}
	re.SetRowsi(result)
	return re, nil
}
