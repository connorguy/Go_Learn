// Database server
// 1. run a server that is accessible on http://localhost:4000/
// 2. receive a request on .../set?somekey=somevalue store the passed key and value in memory
// 3. receives a request on .../get?key=somekey it should return the value stored at somekey

import(
    "net/http"
	"fmt"
    "os"
)

const (
	set = "set"
	getKey = "get"
)

inMemoryKeyValueStore := make(map[string]string)

func logAnErrorAndExit(err string){
	log.Error("Error: %s", err)
	os.Exit(-1)
}

func set(w http.ResponseWriter, req *http.Request) {
	if err = req.ParseForm(); err != nil {
		logAnErrorAndExit("Error parsing set")
	}
	reqKeyValueMap := req.Form.Get(set)
	if len(reqKeyValueMap) !=1{
		logAnErrorAndExit("invalid request for set.")
	}
	key := 1
	value :=1

	// Store the value now
	inMemoryKeyValueStore[key] = value
}

func get()(w http.ResponseWriter, req *http.Request) {
	
}


func main(){
	http.HandleFunc("/set", set)
	http.HandleFunc("/get", get)


	http.ListenAndServe(":4000", nil)
}