package ge

import (
	"database/sql"
	"fmt"
	"math"
	"testing"
	"time"
)

var (
	a1 interface{}
	a2 interface{}
	a3 interface{}
	a4 interface{}
	a5 interface{}
)

func TestColumNames(t *testing.T) {
	coltypeStandare := []string{"BOOLEAN", "TIMESTAMP", "REAL", "INTEGER", "VARCHAR"}
	tempQuery := "SHOW COLUMNS FROM testdata"
	db, err := sql.Open("elastic", "localhost:9200")
	defer db.Close()
	if err != nil {
		t.Error(err)
	}
	rows, err := db.Query(tempQuery)
	if err != nil {
		t.Error(err)
	}
	i := 0
	for rows.Next() {
		if err := rows.Scan(&a1, &a2); err != nil {
			t.Error(err)
		}

		if a2 != coltypeStandare[i] {
			t.Error("Excepted: ", coltypeStandare[i])
		}
		i++
		fmt.Println(a1, a2)
	}

}

func TestValues(t *testing.T) {
	tempQuery1 := "SELECT * FROM testdata"
	db, err := sql.Open("elastic", "localhost:9200")
	defer db.Close()
	if err != nil {
		t.Error(err)
	}
	rows, err := db.Query(tempQuery1)
	if err != nil {
		t.Error(err)
	}
	for rows.Next() {

		if err := rows.Scan(&a1, &a2, &a3, &a4, &a5); err != nil {
			t.Error(err)
		}
		if a1 != true {
			t.Error("Excepted: true got: ", a1)
		}

		if tempa2, err := time.Parse("2006-01-02 15:04:05 ", "2009-11-15 14:12:12"); err != nil {
			t.Error(err)
		} else {
			if a2 != tempa2 {
				t.Error("Excepted: 2009-11-15T14:12:12 , got: ", a2)
			}
		}

		if toFixed(a3.(float64), 7) != 3.1415927 {
			t.Error("Excepted: 3.1415927, got: ", toFixed(a3.(float64), 7))
		}

		if a5 != "Hello World!" {
			t.Error("Excepted: Hello World! got: ", a5)
		}
		fmt.Println(a1, a2, toFixed(a3.(float64), 7), a4, a5)
	}
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}
