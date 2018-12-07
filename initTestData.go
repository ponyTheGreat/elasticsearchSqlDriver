package ge

import (
	"bytes"
	"net/http"
	"time"
)

//create a new index "testdata" for testing

func testdataCreat() {
	createQuery := `{
	"mappings":{
		"test1":{
		  "properties":{
			  "intTest":{
				  "type":"integer"
			  },
			  "stringTest":{
				  "type":"text"
				},
			  "floatTest":{
				  "type":"float"
				},
				"date":{
					"type":"date"
			 
			  },
			  "booleanTest":{
					"type":"boolean"
			  }
			}		
		 }
	}
}`

	request := []byte(createQuery)
	requestReader := bytes.NewReader(request)
	url := "http://localhost:9200/testdata"
	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, requestReader)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}
}

const row1 string = `{
	"intTest":1,
	"stringTest":"Hello World!",
	"floatTest": 3.14159265354,
	"date":"2009-11-15T14:12:12",
	"booleanTest": true
	}
	`

func insertDoc1() {
	request := []byte(row1)
	requestReader := bytes.NewReader(request)
	url := "http://localhost:9200/testdata/test1/"
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, requestReader)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}

	//time.Sleep(time.Duration(1) * time.Second)

}

func initTestData() {
	testdataCreat()
	insertDoc1()
	insertDoc1()
	insertDoc1()
	insertDoc1()
	insertDoc1()
	insertDoc1()
	insertDoc1()
	time.Sleep(time.Duration(1) * time.Second)
}

func clearTestData() {
	url := "http://localhost:9200/testdata"
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		panic(err)
	}
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}

}

/*
func getTestDb() *sql.DB {
	db, err := sql.Open("elastic", "localhost:9200")
	if err != nil {
		panic(err)
	}
	return db
}
*/
func getStmt(query, url string) *Stmt {
	stmt := Stmt{
		Method:   "POST",
		SQLQuery: query,
		URL:      url,
	}
	return &stmt
}
