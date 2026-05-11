package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	Mad_lib()
}

func Mad_lib() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your name? ")
	name, _ := reader.ReadString('\n')
	fmt.Print("What is Her name? ")
	herName, _ := reader.ReadString('\n')
	fmt.Print("What do you want her to be? ")
	whatIsTheRelationship, _ := reader.ReadString('\n')
	fmt.Print("Make a prayer : ")
	prayer, _ slice[rune]:= reader.ReadString('\n')
	fmt.Printf("I'm %s. I want %s to be with me as %s. Oneday, %s. :)", name, herName, whatIsTheRelationship, prayer)
}
