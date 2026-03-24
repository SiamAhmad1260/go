package main

import (
	"fmt"
	"unicode/utf8"
)

func CountNumberOfCharactors() {
	var str string
	fmt.Print("what is the string input? ")
	fmt.Scan(&str)
	fmt.Printf("%s has %d charactors.", str, utf8.RuneCountInString(str))
}

func Challenge_1() {
	var str string
	fmt.Print("What is the string? ")
	fmt.Scan(&str)
	if str == "" {
		panic("must have a proper input!")
	} else {
		fmt.Printf("%s contains %d number of charactors.", str, utf8.RuneCountInString(str))
	}
}
