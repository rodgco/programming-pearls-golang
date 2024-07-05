package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/rodgco/bigbitvector"
)

func fromRandom(bv *bigbitvector.BigBitVector) {
	fmt.Println("fromRandom")
	for i := 0; i < 10_000_000; i++ {
		n := rand.Intn(10_000_000)
		err := bv.SetBit(n)

		if err != nil {
			log.Fatal(err)
		}
	}
}

func fromFile(bv *bigbitvector.BigBitVector) {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			log.Fatal(err)
		}
		}(f)

	s := bufio.NewScanner(f)
	for s.Scan() {
		str := s.Text()
		n, _ := strconv.Atoi(str)
		err := bv.SetBit(n)

		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	bv := bigbitvector.New(10_000_000)

	start := time.Now()
	// fromRandom(vec)
	fromFile(bv)

	t1 := time.Now()

	var count int
	for i := 0; i < bv.Size(); i++ {
		test, err := bv.IsBitSet(i)

		if err != nil {
			log.Fatal(err)
		}

		if test {
			count++
			// fmt.Println(i)
		}
	}
	end := time.Now()
	fmt.Println("Start....:", start.Format("15:04:05.000"))
	fmt.Println("Random...:", t1.Format("15:04:05.000"), t1.Sub(start).Milliseconds())
	fmt.Println("End......:", end.Format("15:04:05.000"), end.Sub(t1).Milliseconds())
}

