package ge

import (
	"testing"
)

func TestPrepare(t *testing.T) {
	con := &conn{
		host:    "localhost",
		portnum: "9200",
	}

	stmt, err := con.Prepare("SELECT * FROM people")
	if err != nil {
		t.Error("get error: ", err)
	}
	if stmt.(*Stmt).SQLQuery != "SELECT * FROM people" {
		t.Error("Excepted: \"SELECT * FROM people\" got:", stmt.(*Stmt).SQLQuery)
	}
	if stmt.(*Stmt).URL != "http://localhost:9200" {
		t.Error("Excepted: \"http://localhost:9200\" got:", stmt.(*Stmt).URL)
	}
}

func TestConnQuery(t *testing.T) {
	initTestData()
	con := &conn{
		host:    "localhost",
		portnum: "9200",
	}
	result, err := con.Query("SELECT * FROM testdata")
	if err != nil {
		t.Error(err)
	}
	if result == nil {
		t.Error("no return")
	}
}
