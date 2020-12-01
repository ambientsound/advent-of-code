package lib_test

import (
	"io/ioutil"
	"testing"

	"github.com/ambientsound/aoc/lib"
)

func BenchmarkSum2(b *testing.B) {
	f, err := ioutil.ReadFile("input.txt")
	ints := lib.Readints(f)

	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		lib.Sums2(2020, 2, []int{}, ints)
	}
}

func BenchmarkSum3(b *testing.B) {
	f, err := ioutil.ReadFile("input.txt")
	ints := lib.Readints(f)

	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		lib.Sums2(2020, 3, []int{}, ints)
	}
}
