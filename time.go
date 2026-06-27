package main

import (
	"fmt"
	"time"
)

func main() {
	current := time.Now()
	year := current.Year()
	fmt.Println(year)
}
