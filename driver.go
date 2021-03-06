package ge

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strings"
)

//Elastic for driver
type Elastic struct{}

//Open for driver return a connection of DB
func (d *Elastic) Open(dsn string) (driver.Conn, error) {
	content := strings.Split(dsn, ":")
	hostName := content[0]
	port := content[1]
	dc := &conn{
		host:    hostName,
		portnum: port,
	}
	return dc, nil
}

func init() {
	sql.Register("elastic", &Elastic{})
	fmt.Println("register successfully")
}
