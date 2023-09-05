package main

import (
	"fmt"
	"math"
	// "github.com/gifted/go/strutil"
)

func conditional(nameArr []string) {
	if len(nameArr) < 2 {
		fmt.Println("Scenario 1", nameArr)
	} else {
		fmt.Println(nameArr[0:1])
	}

	// fmt.Println(strutil.ReverseString(nameArr[0]))
}

func main() {
	pow := math.Floor(1.7)
	nameArr := []string{"GiFTED!"}
	fmt.Println("Hello World!", pow)

	conditional(nameArr)
}
