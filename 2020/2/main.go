package main

import (
	"fmt"
	"io/ioutil"

	lib2 "github.com/ambientsound/aoc/2/lib"
)

func main() {
	raw, _ := ioutil.ReadFile("/Users/kimt/src/aoc/2020/2/input.txt")
	policies := lib2.ReadFile(raw)
	for _,policy := range policies{
		if policy.Valid2() {
			fmt.Printf("valid: %s\n", policy.Password)
		} else {
			fmt.Printf("invalid: %s\n", policy.Password)
		}
	}
}
