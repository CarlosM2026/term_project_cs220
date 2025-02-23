// Author: Carlos Moncada Soto, Abby Baron, Ken Lam
// Program Name: coding.go
// Date: 02/23/2025 
// Description: reads input and gives output based on what it was given
// Usage: go run coding.go      gives user input

// Reference: https://go.dev/doc/code for Go setup and terminology insights

// Packages group related functions and consist of all files within the same directory
package main

// The fmt package provides functions for text formatting, including console output.
// It is part of Go's standard library, included with the Go installation.
import "fmt"

// Object for creating user inputted data
type Person struct {
    FullName string
    Age  int
    Hobby string
}

// The main function is the entry point of the program and runs automatically when the main package is executed.
// Main function currently asks user for space-seperated input and shows your input
func main() {
	Carlos := Person{FullName: "Carlos Moncada", Age: 20, Hobby: "Soccer"}

	fmt.Print("Enter your name, age, and hobby seperated with a space: ")

    fmt.Scan(&Carlos.FullName, &Carlos.Age, &Carlos.Hobby)

	fmt.Println("Your name is:", Carlos.FullName, "\nYour age is:", Carlos.Age, "\nYour hobby is:", Carlos.Hobby)
}
