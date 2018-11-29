package main

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {
	log.SetFlags(log.Llongfile)
	//query := "SELECT * FROM library"
	tempQuery := fmt.Sprintf("SELECT * FROM %s ORDER BY id ASC", "people")
	fmt.Println(tempQuery)
	//fmt.Println(tempQuery)
	//q := "SELECT * FROM library"
	db, err := sql.Open("ESDriver", "localhost:9200")
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
}
