package ge

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strings"
)

//ESDriver struct
type ESDriver struct{}

//Open for driver return a connection of DB
func (d *ESDriver) Open(dsn string) (driver.Conn, error) {
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
	sql.Register("ESDriver", &ESDriver{})
	fmt.Println("register successfully")
}
