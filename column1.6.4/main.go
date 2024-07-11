package main

import (
	"fmt"
	"math/rand"

	bbv "github.com/rodgco/bigbitvector"
)

const (
	n = 10_000_000
	k = 100_000
)

func main() {
	// Create a new bit vector
	bv := bbv.New(n)
	bv.SetAll()
	for c := 0; c < k; c++ {
		r := rand.Intn(bv.Count())+1
		index, err := bv.FindNthSet(r)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(index)
		bv.Unset(index)
	}
}
