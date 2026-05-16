package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Mad_lib() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your name? ")
	name, _ := reader.ReadString('\n')
	fmt.Print("What is Her name? ")
	herName, _ := reader.ReadString('\n')
	fmt.Print("What do you want her to be? ")
	whatIsTheRelationship, _ := reader.ReadString('\n')
	fmt.Print("Make a prayer : ")
	prayer, _ := reader.ReadString('\n')

	cleanName := strings.TrimRight(name, "\n\r")
	cleanHerName := strings.TrimRight(herName, "\n\r")
	cleanRelationship := strings.TrimRight(whatIsTheRelationship, "\n\r")
	cleanPrayer := strings.TrimRight(prayer, "\n\r")
	fmt.Print("I'm ", cleanName, ". I want ", cleanHerName, " to be with me as ", cleanRelationship, ". Oneday, ", cleanPrayer, ". :)")
}
