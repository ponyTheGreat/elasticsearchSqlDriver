package ge

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type column struct {
	Name    string
	Coltype string `json:"type"`
}

// DefaultFetchSize define the page size
const DefaultFetchSize int = 5

//querywrapedQuery is http request body
func sendHTTPRequestQuery(request []byte, httpMethod, URL string) (*http.Response, error) {
	requestReader := bytes.NewReader(request)
	wrapURL := fmt.Sprintf("%s/_xpack/sql?format=json", URL)
	res, err := http.NewRequest(httpMethod, wrapURL, requestReader)
	if err != nil {
		return nil, err
	}
	res.Header.Add("Content-Type", "application/json")
	client1 := &http.Client{}
	response, err := client1.Do(res)
	if err != nil {
		return nil, err
	}
	return response, nil
}

//last true:this is the last page, false:not the last page
func sendHTTPRequestCursor(cursor, URL string, last bool) (*http.Response, error) {
	client := &http.Client{}
	cr := cursorRequest{
		Cursor: cursor,
	}
	crj, err := json.Marshal(cr)
	if err != nil {
		return nil, err
	}
	cursorReader := bytes.NewReader(crj)

	var wrapURL string
	// not last page
	if !last {
		wrapURL = fmt.Sprintf("%s/_xpack/sql?format=json", URL)
	} else {
		wrapURL = fmt.Sprintf("%s/_xpack/sql/close", URL)
	}
	request, err := http.NewRequest("POST", wrapURL, cursorReader)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func readHTTPResponse(res *http.Response) ([]byte, error) {
	result, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	return result, nil
}

type requestBody struct {
	Query     string `json:"query"`
	FetchSize int    `json:"fetch_size"`
}
type cursorRequest struct {
	Cursor string `json:"cursor"`
}

//outType define the recieve type of json in http reaspons
//Rows [][] get every element in every row
type outType struct {
	Columns []column
	Rows    [][]interface{}
	Cursor  string `json:"cursor"`
}

func getRows(columns []column, rows [][]interface{}, Cursor string, URL string) Rows {
	re := Rows{
		RowsContent:    rows,
		ColumnsContent: columns,
		Cur:            0,
		Closed:         false,
		cursor:         Cursor,
		url:            URL,
	}
	return re
}

func jsonDecode(sample []byte) (*outType, error) {

	var c outType
	if err := json.Unmarshal(sample, &c); err != nil {
		return nil, err
	}
	return &c, nil

}

//convert datatype to driver.Value(int64,float64,bool,[]byte,string,time.Time)
func convertToValue(data interface{}, typeName string) (driver.Value, error) {
	if !driver.IsValue(data) {
		return nil, errors.New("Not a type")
	}
	if typeName == "date" {
		var t time.Time
		tem := fmt.Sprintf("\"%s\"", data.(string))
		err := json.Unmarshal([]byte(tem), &t)
		if err != nil {
			return nil, err
		}
		return t, nil
	}
	return data, nil
}


