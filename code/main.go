package main

import (
	"fmt"
	"io/ioutil"
)

var Words string
var AllWords string

func main() {
	data, err := ioutil.ReadFile("motlist.txt")
	if err != nil {
		fmt.Println("File reading error", err)
	}
	lines := string(data)
	AllWords = lines
	Hangman()
}
