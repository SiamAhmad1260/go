package main

import (
	"fmt"
	"math/rand"
)

func main() {
	Challenge1()
}

func Saying_Hello() {
	var name string
	fmt.Print("What is your name? ")
	fmt.Scan(&name)
	fmt.Printf("Hello, %s, nice to meet you!", name)
}

func Challenge1() {
	var name string
	fmt.Print("What is your name? ")
	fmt.Scan(&name)
	greeting := rand.Intn(7)
	if greeting == 0 {
		fmt.Printf("Hello, %s", name)
	} else if greeting == 1 {
		fmt.Printf("Greetings, %s", name)
	} else if greeting == 2 {
		fmt.Printf("How's it going, %s", name)
	} else if greeting == 3 {
		fmt.Printf("Meow meow %s", name)
	} else if greeting == 4 {
		fmt.Printf("Assalamu Alaikum, %s", name)
	} else if greeting == 5 {
		fmt.Print("Ja, Ja, Ja, Jaann!")
	} else {
		fmt.Print("I miss you!")
	}
}
