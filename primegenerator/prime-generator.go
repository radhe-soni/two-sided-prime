package main

import (
	"fmt"
	"log"
	"runtime"

	primetest "github.com/radhe-soni/two-sided-prime/primetester/api"
)

func main() {
	log.Println("start")
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i := int64(30); i < 1<<20; i++ {
		go func(j int64) {
			if primetest.IsTwoSidedPrime(j) {
				fmt.Println(j)
			}
		}(i)

	}
	log.Println("end")
}
