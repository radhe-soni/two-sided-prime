package main

import (
	"fmt"
	"log"

	primetester "github.com/radhe-soni/two-sided-prime/primetester/api"
)

func main() {
	log.Println("###### A Program which tells if an integer is prime ######")
	log.Print("Enter an integer to continue...\t")
	var x int64
	fmt.Scan(&x)

	result := primetester.IsTwoSidedPrime(x)

	log.Printf("user input is %v\n", result)
}
