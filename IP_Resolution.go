package main

import (
	"fmt"
	"net"
)

// Given a web address resolve the ip associated with it.
func main() {
	var nameToLookUp string
	fmt.Scan(&nameToLookUp)

	ip, err := net.LookupIP(nameToLookUp)

	if err == nil {
		fmt.Println(ip)
	}
}
