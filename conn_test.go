package ge

import (
	"testing"
)

func TestScussPrepare(t *testing.T) {
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

