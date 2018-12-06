package ge

import (
	"database/sql"
	"fmt"
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
	initTestData()
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
	clearTestData()
}

func TestValues(t *testing.T) {
	initTestData()
	/*
		requestBody := `{
				"query":"SELECT * FROM testdata"
				}`
		reqR := bytes.NewReader([]byte(requestBody))
		url := "http://localhost:9200/_xpack/sql?format=json"
		request, err := http.NewRequest("POST", url, reqR)
		if err != nil {
			panic(err)
		}
		request.Header.Add("Content-Type", "application/json")
		client := http.Client{}

		res, err := client.Do(request)
		if err != nil {
			panic(err)
		}
		resp, err := readHTTPResponse(res)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(resp))
	*/
	db, err := sql.Open("elastic", "localhost:9200")
	if err != nil {
		t.Error(err)
	}
	//fmt.Println(db)
	defer db.Close()
	rows, err := db.Query("SELECT * FROM testdata")
	if err != nil {
		t.Error(err)
	}
	//fmt.Println(rows)
	for rows.Next() {
		//time.Sleep(time.Duration(1) * time.Second)
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

		if a5 != "Hello World!" {
			t.Error("Excepted: Hello World! got: ", a5)
		}
		fmt.Println(a1, a2, a3, a4, a5)
	}

	clearTestData()
}

//1. if the last digit <5, then dont keep it
//2. if the last digit ==5 and there is none zero digits behind ,then see if the it is odd ，digit up, if it is even dont keep it
//3. if the last digit ==5 and there is none zero digit > 5, digit up
