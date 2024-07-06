package util

import "fmt"

var debugString string

// Debug prints a string to the bottom of the game window.
func Debug(a any) {
	debugString += fmt.Sprint(a) + "; "
	fmt.Print(debugString)
}
