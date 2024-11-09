package main

import (
	"fmt"
)

func Visokosnygod(x int) string {
	if (x%400 == 0) && (x%100 == 0) {
		return "true"
	} else if (x%4 == 0) && (x%100 != 0) {
		return "true"
	} else {
		return "false"
	}
}

var a int

func main() {
	fmt.Scan(&a)
	fmt.Print(Visokosnygod(a))
}
