package main

import (
	"fmt"
	"io/ioutil"

	"github.com/ambientsound/aoc/lib"
)

func main() {
	f, err := ioutil.ReadFile("input.txt")
	ints := lib.Readints(f)

	if err != nil {
		panic(err)
	}

	result, _ := lib.Sums2(2020, 2, []int{}, ints)
	fmt.Printf("%d x %d = %d\n", result[0], result[1], lib.Multiply(result))

	result, _ = lib.Sums2(2020, 3, []int{}, ints)
	fmt.Printf("%d x %d x %d = %d\n", result[0], result[1], result[2], lib.Multiply(result))
}
