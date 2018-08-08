package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// getRESTCall is a more generic implementation of http request for reusability.
// Gets a http response from RESTful api (or any json formatted response). Body
// of the response is put into a []byte and then unmarshaled into a generic
// interface which is returned to the calling function. Referenced:
// https://medium.com/@IndianGuru/consuming-json-apis-with-go-d711efc1dcf9
func getRESTCall(url string) interface{} {
	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil
	}

	// For control over HTTP client headers, redirect policy, and other
	// settings, create a Client A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client Do sends an HTTP request and returns an
	// HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("http client do error: ", err)
		return nil
	}

	// Callers should close resp.Body when done reading from it Defer the
	// closing of the body
	defer resp.Body.Close()

	// Put body of the response into a []byte
	body, err := ioutil.ReadAll(resp.Body)

	// Put body of the json into an interface var
	var unstructuredJSON interface{}
	err = json.Unmarshal(body, &unstructuredJSON)
	if err != nil {
		log.Fatal("Json unmarshal error: ", err)
		return nil
	}

	return unstructuredJSON
}

// jsonMapper will print out the map of json data collected.
func jsonMapper(unstructuredJSON interface{}) {
	jsonBody := unstructuredJSON.(map[string]interface{})

	for k, v := range jsonBody {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, ":", vv)
		case float64:
			fmt.Println(k, ":", vv)
		case []interface{}:
			fmt.Println(k, ":")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, " - is of a type I don't know how to handle")
		}
	}
}

func main() {

	// Get input from console reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter json api http: ") input, _ := reader.ReadString('\n')

	// Fot the purpose of testing a test api
	input := "https://jsonplaceholder.typicode.com/posts/1/comments"

	unstructuredJSON := getRESTCall(input)
	jsonMapper(unstructuredJSON)

}
