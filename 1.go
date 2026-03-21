package main

import (
	"fmt"
	"math/rand"
)

func main() {
	Challenge_1_Switches()
}

func Saying_Hello() {
	var name string
	fmt.Print("What is your name? ")
	fmt.Scan(&name)
	fmt.Printf("Hello, %s, nice to meet you!", name)
}

func Challenge_1() {
	var name string
	fmt.Print("What is your name? ")
	fmt.Scan(&name)
	greeting := rand.Intn(7)
	if greeting == 0 {
		fmt.Printf("Hello, %s", name)
	} else if greeting == 1 {
		fmt.Printf("Greetings, %s", name)
	} else if greeting == 2 {
		fmt.Printf("How's it going %s?", name)
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

func Challenge_1_Switches() {
	var name string
	fmt.Print("What is your name? ")
	fmt.Scan(&name)
	greeting := rand.Intn(7)
	switch {
	case greeting == 0:
		{
			fmt.Printf("Hello, %s", name)
		}
	case greeting == 1:
		{
			fmt.Printf("Greetings, %s", name)
		}
	case greeting == 2:
		{
			fmt.Printf("How's it going %s?", name)
		}
	case greeting == 3:
		{
			fmt.Printf("Assalamu Alaikum, %s", name)
		}
	case greeting == 4:
		{
			fmt.Printf("meow meow, %s", name)
		}
	case greeting == 5:
		{
			fmt.Printf("Assalamu alaikum, ...")
		}
	case greeting == 6:
		{
			fmt.Printf("How are you, ...")
		}
	}
}
