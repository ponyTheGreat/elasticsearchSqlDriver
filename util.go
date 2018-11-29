package ge

import (
	"bufio"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type column struct {
	Name    string
	Coltype string `json:"type"`
}

// DefaultFetchSize define the page size
const DefaultFetchSize int = 5

//querywrapedQuery is http request body
func sendHTTPRequestQuery(wrapedQuery, httpMethod, URL string) (*http.Response, error) {
	request := stringToIR(wrapedQuery)
	wrapURL := fmt.Sprintf("%s/_xpack/sql?format=json", URL)
	res, err := http.NewRequest(httpMethod, wrapURL, request)
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
	requestBody := fmt.Sprintf("{\n\"cursor\":\"%s\"\n}", cursor)
	var wrapURL string
	// not last page
	if !last {
		wrapURL = fmt.Sprintf("%s/_xpack/sql?format=json", URL)
	} else {
		wrapURL = fmt.Sprintf("%s/_xpack/sql/close", URL)
	}
	requestBodyReader := stringToIR(requestBody)
	request, err := http.NewRequest("POST", wrapURL, requestBodyReader)
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

//convert string to io.Reader
func stringToIR(str string) io.Reader {
	str1 := strings.NewReader(str)
	result := bufio.NewReader(str1)
	return result
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

//MakeLines split string by lines
func MakeLines(s string) []string {
	lines := strings.Split(s, "\n")
	return lines
}

//SubString for get a substring start with c1, end with c2 from a string
func SubString(src string, c1, c2 rune) (string, error) {
	i := -1
	j := -1
	for index, cont := range src {
		if cont == c1 {
			i = index
		}
		if cont == c2 {
			j = index
		}
	}
	if i == -1 || j == -1 {
		return "", errors.New("no char found")
	}
	if i > j {
		return "", errors.New("no substring")
	}
	return src[i:j], nil
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
	return data.(driver.Value), nil
}
