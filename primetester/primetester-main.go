package primetester

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func main() {
	log.Println("###### A Program which tells if an integer is prime ######")
	log.Println("Enter an integer to continue...")
	var x int64
	fmt.Scan(&x)
	numberOfDigits := len(strconv.FormatInt(x, 10))
	result := IsPrime(x)
	for i := 1; i < numberOfDigits && result; i++ {
		divisor := int64(math.Pow10(i))
		multiplier := int64(math.Pow10(i - 2))
		firstPart := (x % divisor) * multiplier
		secPart := x / divisor
		fmt.Println("firstPart", firstPart)
		fmt.Println("secPart", secPart)
	}
	log.Printf("user input is %v\n", result)
}

/*
* Check if an integer is prime.
 */
func IsPrime(test int64) bool {
	result := true
	if test < 2 {
		result = false
	} else {
		sqrt := math.Sqrt(float64(test))
		for i := int64(3); i < int64(sqrt); i++ {
			if test%i == 0 {
				result = false
			}
		}
	}

	return result
}
