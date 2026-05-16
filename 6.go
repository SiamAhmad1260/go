package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	current := time.Now()

	yearNow := current.Year()

	fmt.Print("what year are you going to retire : ")
	input, _ := reader.ReadString('\n')

	cleanInput := strings.TrimRight(input, "\n\r")
	retireTime, _ := strconv.Atoi(cleanInput)

	timeLeft := (retireTime - yearNow)

	fmt.Printf("The time left is : %v", timeLeft)
}
