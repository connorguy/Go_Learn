package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Utility program for files. Files are required to be in the same directory.
func main() {
	// Get file name from user
	fmt.Print("Please enter file name in this dir: ")
	var fileName string
	fmt.Scan(&fileName)

	// Open file and defer closure
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Execute command while user still wants
	newCommand := true
	for newCommand {
		fmt.Println()
		printOptions()
		newCommand = getAction(file)
	}

}

// getAction is called after the prompt for imput is given and the user enters
// what command to issue. Function will parse this input and execute appropriate
// command. Returns true unless user enters exit command.
func getAction(file *os.File) bool {
	// Get user input
	var action string
	fmt.Scan(&action)
	// Normalize string to uppercase
	action = strings.ToUpper(action)

	// Call the appropriate command. If invalid returns to calling function and
	// returns true for further commands. If E returns false implying theres no
	// further command.
	switch action {
	case "RL":
		readByLine(file)
	case "S":
		printStats(file)
	case "E":
		return false
	default:
		fmt.Println("No command found")
	}

	return true
}

// printOptions prints the list of available commands.
func printOptions() {
	fmt.Println("==================")
	fmt.Println("So what should I do?")
	//fmt.Println("R: read all contents")
	fmt.Println("RL: read by line")
	fmt.Println("S: output system stats on file")
	fmt.Println("E: Exit")
	fmt.Println("==================")

}

// printStats given a file pointer will print the name, size, system mode, and
// system stat on a file.
func printStats(file *os.File) {
	// Get file stats
	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("=== File Stats ===")
	fmt.Println("Name: ", stat.Name())
	fmt.Println("Size: ", stat.Size())
	fmt.Println("Mod: ", stat.Mode())
	fmt.Println("Sys: ", stat.Sys())
	fmt.Println("==================")
}

// readByLine takes a file pointer and outputs one line from the file at a time.
// Will ask user if they would like to continue after each line.
func readByLine(file *os.File) {
	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())

		if !shouldContinue() {
			break
		}

	}
}

// shouldContinue is a utility function for a yes/no prompt. Returns true for
// yes, false for no.
func shouldContinue() bool {
	var cont string
	fmt.Println("\nContinue? Y/N")
	fmt.Scan(&cont)
	cont = strings.ToLower(cont)
	return strings.Contains(cont, "y")
}
