package main

import (
	"fmt"
	"strings"
)

func main() {
	var message = "I Love Golang"
	fmt.Scanln(&message)
	// split string from space " "
	splittedString := strings.Split(message, " ")

	fmt.Println(splittedString)
}
