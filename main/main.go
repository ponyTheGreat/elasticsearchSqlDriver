package main

import (
	"bytes"
	"fmt"
	"net/http"

	_ "github.com/lib/ge"
)

// main func used to create a new index for testing
func main() {
	/*
		log.SetFlags(log.Llongfile)
		//query := "SELECT * FROM library"
		tempQuery := fmt.Sprintf("SELECT * FROM %s ORDER BY id ASC", "people")
		fmt.Println(tempQuery)
		//fmt.Println(tempQuery)
		//q := "SELECT * FROM library"
		db, err := sql.Open("elastic", "localhost:9200")
		if err != nil {
			panic(err)
		}
		defer db.Close()

		rows, err := db.Query(tempQuery)
		//fmt.Println("------------")
		//col, _ := rows.Columns()
		//len := len(col)
		//fmt.Println(len)
		if err != nil {
			panic(err)
		}

		defer rows.Close()

		var (
			a1 interface{}
			a2 interface{}
			//a3 interface{}
			//a4 interface{}
		)

		var tt []interface{}
		for i := 0; i < 4; i++ {
			var tem interface{}
			tt = append(tt, tem)
		}
		for rows.Next() {

			if err := rows.Scan(&a1, &a2); err != nil {
				log.Fatal(err)
			}
			fmt.Println(a1, a2)
		}
	*/
	//create test index and insert data
	testdataCreat()
	insertDoc1()
}

//create a new index "testdata" for testing
const createQuery string = `{
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

func testdataCreat() {
	request := []byte(createQuery)
	requestReader := bytes.NewReader(request)
	url := "http://localhost:9200/testdata"
	client := &http.Client{}
	res, err := http.NewRequest("PUT", url, requestReader)
	if err != nil {
		panic(err)
	}
	res.Header.Add("Content-Type", "application/json")
	result, err := client.Do(res)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
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
	url := "http://localhost:9200/testdata/test1"
	client := &http.Client{}
	res, err := http.NewRequest("POST", url, requestReader)
	if err != nil {
		panic(err)
	}
	res.Header.Add("Content-Type", "application/json")
	result, err := client.Do(res)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

}
