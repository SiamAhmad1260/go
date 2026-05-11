package main

import (
	"fmt"
	"math/rand/v2"
)

type quotes struct {
	name  string
	quote map[int]string
}

// must increase the functionality and make more quotes
func PrintQuotes() {
	first := quotes{name: "Joe Biden",
		quote: map[int]string{
			1: "What?",
			2: "The next president of the United States - Barak America",
			3: "Stand up, Chuck, let'em see ya.",
			4: "You cannot go to a 7-11 or a Dunkin' Donuts unless you have a slight Indian accent.... I'm not joking.",
			5: "The US is a nation that can be discribed in a single word - \"asufutimaehaehfutbw\"",
		},
	}
	quoteNumber := rand.IntN(5)
	quot := first.quote[quoteNumber]
	name := first.name
	fmt.Printf("Whats the quote?\n\"%s\" by %s", quot, name)
}
