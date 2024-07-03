package main

var debugString string

// Debug prints a string to the bottom of the game window.
func Debug(s string) {
	debugString += s + "; "
	print(debugString)
}
