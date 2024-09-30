package main

import (
	"fmt"
	"io/ioutil"
)

var Words string
var List []string

func main() {
	data, err := ioutil.ReadFile("motlist.txt")
	if err != nil {
		fmt.Println("File reading error", err)
	}
	lines := string(data)

	for _, line := range lines {
		if line != '\n' && line != '\r' {
			Words += string(line)
		} else {
			List = append(List, Words)
			Words = ""
		}
	}
	fmt.Println(List)
	Hangman()
}
