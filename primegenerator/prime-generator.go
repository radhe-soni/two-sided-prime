package main

import (
	"fmt"

	primetester "github.com/radhe-soni/two-sided-prime/primetester/api"
)

func main() {
	for i := int64(30); i < 1<<15; i++ {
		if primetester.IsPrime(i) {
			fmt.Println(primetester.IsPrime(i))
			fmt.Println(i)
		}
	}
}
