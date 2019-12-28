package primetester

import (
	"fmt"
	"math"
	"strconv"
)

//IsPrime : Check if an integer is prime.
func IsPrime(test int64) bool {
	result := true
	if test <= 1 {
		result = false
	} else {
		sqrt := math.Sqrt(float64(test))
		for i := int64(2); i <= int64(sqrt); i++ {
			if test%i == 0 {
				result = false
			}
		}
	}

	return result
}

//IsTwoSidedPrime : Checks if an integer is two sided prime
func IsTwoSidedPrime(test int64) bool {
	numberOfDigits := len(strconv.FormatInt(test, 10))
	result := IsPrime(test)
	for i := 1; i < numberOfDigits && result; i++ {
		divisor := int64(math.Pow10(i))
		multiplier := int64(math.Pow10(i - 1))
		fmt.Println("divisor", divisor, "\tmultiplier", multiplier)
		firstPart := (test % divisor) * multiplier
		secPart := test / divisor
		fmt.Println("firstPart", firstPart)
		fmt.Println("secPart", secPart)
	}
	return result
}
