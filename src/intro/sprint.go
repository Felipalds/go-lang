// Golang program to illustrate the usage of
// fmt.Sprint() function

// Including the main package
package main

// Importing fmt, io and os
import (
	"fmt"
	"io"
	"os"
)

// Calling main
func main() {

	// Declaring some const variables
	const name, dept = "GeeksforGeeks", "CS"

	// Calling Sprint() function
	s := fmt.Sprint(name, " is a ", dept, " Portal.\n")

	// Calling WriteString() function to write the
	// contents of the string "s" to "os.Stdout"
	io.WriteString(os.Stdout, s)

}
