package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Given the a set of integers, return their indexed values index. ie arr[n] =
// x, return arr[x] first input is number of integers with a return second input
// is the integers seperated by spaces
func main() {
	// reads in the next line as a string
	numbers := getIntegerArray()

	// Return what the indexed number indexes to
	for i, _ := range numbers {
		indexToGet := numbers[i] - 1
		if indexToGet >= 0 && indexToGet < len(numbers) {
			fmt.Println(numbers[indexToGet])
		}
	}

}

// getIntegerArray Reads line in from stdin and converts to integer array - has
// to be delimiter by spaces.
func getIntegerArray() []int {
	// Get number of integers coming, but not kept since well use the string[]
	// len as the value.
	var numOfNums int
	fmt.Scan(&numOfNums)

	// reads in the next line as a string
	reader := bufio.NewReader(os.Stdin)
	numbers, _ := reader.ReadString('\n') // Note that as far as I could find, strings don't have a capacity cap - besides your available mem.
	// Preproccess string to make sure it ends on a number not a space or '\n' -
	// does not check for EOF.
	numbers = strings.TrimSuffix(numbers, " \n")
	numbers = strings.TrimSuffix(numbers, "\n")
	numbers = strings.TrimSuffix(numbers, " ")

	// Creates an int array and parses string with space delimiter 
	numAsStringArray := strings.Split(numbers, " ")

	// Convert to int and store in array that we return)
	intArray := make([]int, len(numAsStringArray))
	for i, stringNum := range numAsStringArray {
		num, _ := strconv.Atoi(stringNum)
		intArray[i] = num
	}

	return intArray
}
