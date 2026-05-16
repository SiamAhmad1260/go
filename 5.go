package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stringtoformat() {
	var fst, snd string
	fmt.Print("what is the number : ")
	fmt.Scan(&fst)
	fmt.Print("what is the 2nd number : ")
	fmt.Scan(&snd)

	num1, _ := strconv.ParseInt(fst, 10, 32)
	num2, _ := strconv.ParseInt(snd, 10, 32)
	plus, minus := num1+num2, num1-num2
	fmt.Printf("%v + %v = %v\n%v - %v = %v", num1, num2, plus, num1, num2, minus)
}
func challenge5() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("what is the number : ")
	fst, _ := reader.ReadString('\n')
	cleanFst := strings.TrimRight(fst, "\n\r")
	fstNew, err1 := strconv.Atoi(cleanFst)

	fmt.Print("what is the 2nd number : ")
	snd, _ := reader.ReadString('\n')
	cleanSnd := strings.TrimRight(snd, "\n\r")
	sndNew, err2 := strconv.Atoi(cleanSnd)

	plus := fstNew + sndNew
	minus := fstNew - sndNew

	if err1 != nil || err2 != nil {
		fmt.Println("Please input a number")
	} else if fstNew < 0 || sndNew < 0 {
		fmt.Println("No negetive number is accepted")
	} else {
		fmt.Printf("%v + %v = %v\n%v - %v = %v", fstNew, sndNew, plus, fstNew, sndNew, minus)
	}

}
