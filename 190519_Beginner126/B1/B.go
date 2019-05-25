package main

import (
	"fmt"
	"strconv"
	"strings"
)

var S string

func main() {
	fmt.Scan(&S)
	a := strings.Split(S, "")
	b := a[0] + a[1]
	c := a[2] + a[3]
	yy, _ := strconv.Atoi(b)
	mm, _ := strconv.Atoi(c)
	if yy >= 1 && yy <= 12 {
		if mm >= 1 && mm <= 12 {
			fmt.Println("AMBIGUOUS")
		} else {
			fmt.Println("MMYY")
		}
	} else {
		if mm >= 1 && mm <= 12 {
			fmt.Println("YYMM")
		} else {
			fmt.Println("NA")
		}
	}
}
