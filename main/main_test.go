package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFibb(t *testing.T) {
	correctFibb := []int{4181, 6765, 10946, 17711, 28657, 46368, 75025, 121393, 196418, 317811}
	fibb := Fibb{Id: 001}
	fibb.CalcFibb()
	if !assert.Equal(t, correctFibb, fibb.GetFunky()) {
		t.Fatal("He dead")
	}

}
