package main

import (
	"fmt"
)

// Hacker Rank Hash Tables: Ransom Note Make try and make a given note from the
// available words, no duplicates. Output Yes or No if possible.
func main() {

	m := []string{"hello", "no", "yes", "fun"}
	n := []string{"no", "fun"}

	checkMagazine(m, n)

}

func checkMagazine(magazine []string, note []string) {
	stringMap := createMap(magazine)

	for _, word := range note {
		val, exists := stringMap[word]
		// Check if either we ran out of the word or it just never existed
		if val == 0 || !exists {
			fmt.Println("NO")
		}

		if exists {
			val--
			stringMap[word] = val
		}
	}

	fmt.Println("YES")

}

func createMap(m []string) map[string]int {
	var stringMap map[string]int
	// Need to initialize the map otherwise go will write to nil and panic
	stringMap = make(map[string]int)

	// Go through the whole array of strings and add them to the map
	for _, message := range m {
		// Check if we have that string in the map
		val, exists := stringMap[message]

		// if the string exists increment its count and set the value
		if exists {
			val++
			stringMap[message] = val
		} else {
			// if it doesn't exist then we add and set count to one
			stringMap[message] = 1
		}
	}

	return stringMap
}
