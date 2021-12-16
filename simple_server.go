// Database server
// 1. run a server that is accessible on http://localhost:4000/
// 2. receive a request on .../set?somekey=somevalue store the passed key and value in memory
// 3. receives a request on .../get?key=somekey it should return the value stored at somekey

import(
    "net/http"
)

func set(w http.ResponseWriter, req *http.Request) {
	
}

func get()(w http.ResponseWriter, req *http.Request) {
	
}


func main(){
	http.HandleFunc("/set", set)
	http.HandleFunc("/get", get)


	http.ListenAndServe(":4000", nil)
}