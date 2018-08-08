package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Check whether a number is contained within a certain set.
func main() {
	// Initial number were comparing to
	var numberToComp int
	fmt.Scan(&numberToComp)

	// reads in the next line as a string
	numbers := getIntegerSet()
	// note if int is not contained returns the base value of the type ie: false
	contains := numbers[numberToComp]

	if contains {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

// getIntegerSet Reads line in from stdin and converts to integer set - has to
// be delimiter by spaces.
func getIntegerSet() map[int]bool {
	// reads in the next line as a string
	reader := bufio.NewReader(os.Stdin)
	numbers, _ := reader.ReadString('\n')

	// Creates an int array and parses string with space seperation
	numAsStringArray := strings.Split(numbers, " ")
	// Create a set (tech a map to bool)
	intSet := make(map[int]bool)
	for _, stringNum := range numAsStringArray {
		num, _ := strconv.Atoi(stringNum)
		intSet[num] = true
	}

	return intSet
}
