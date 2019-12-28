package primegenerator

import (
	"fmt"

	"github.com/radhe-soni/two-sided-prime/primetester"
)

func main() {
	for i := int64(30); i < 1<<50; i++ {
		if primetester.IsPrime(i) {
			fmt.Println(i)
		}
	}
}
