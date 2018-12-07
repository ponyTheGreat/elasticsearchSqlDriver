package ge

import (
	"testing"
	"time"
)

//test functions of stmt and row

func TestQuery(t *testing.T) {
	initTestData()
	insertDoc1()
	insertDoc1()
	time.Sleep(time.Duration(2) * time.Second)
	stmt := getStmt("SELECT * FROM testdata", "http://localhost:9200")

	rows, err := stmt.Query(nil)
	if err != nil {
		t.Error(err)
	}
	defer rows.Close()
	clearTestData()
}
