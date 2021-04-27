package main

import (
	"math"
	"math/rand"
)

var testSlice = generateRandomIntSlice()

func getSliceCopy() []int {
	newslice := make([]int, len(testSlice))
	copy(newslice, testSlice)
	return newslice
}

func generateRandomIntSlice() []int {
	s := make([]int, 142000)
	for i := range s {
		s[i] = rand.Intn(math.MaxInt32)
	}
	return s
}

type custIntSlice []int

func (cs custIntSlice) Len() int {
	return len(cs)
}

func (cs custIntSlice) Less(i, j int) bool {
	return cs[i] < cs[j]
}

func (cs custIntSlice) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}
