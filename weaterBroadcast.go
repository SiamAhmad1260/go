package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Use pointers as much as you can

type weather struct {
	temperature int
	humidity    float32
	cO2Content  float32
	season      seasons
}

type seasons int
type inputs any
type input interface{
	input() 
}


const (
	summer seasons = iota
	autumn
	winter
	spring
)

func input() int {
	reader := bufio.NewReader(os.Stdin)
	d, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Input Problem\n")
	}
	cleanD := strings.TrimRight(d, "\n")
	D := conv.Atoi(cleanD)
	return D
}

func (w *weather) environmentHealth(temperature int, humidity float32, cO2Content float32) float32 {
	breathability := float32(w.temperature) * w.humidity * w.cO2Content
	return breathability
}

func main() {
	var w weather
	fmt.Println("whats todays tempurature: ")
	w.temperature := input()
	temp := &w.temperature

	fmt.Println("whats is the humidity")
	w.humidity := input()
	hum := &w.humidity
	fmt.Println("co2 count ")
	w.cO2Content := input()

	num := breathability(temp)

}
