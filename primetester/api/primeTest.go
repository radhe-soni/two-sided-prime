package primetest

import (
	"math"
	"strconv"
)

//AsyncIsPrime : Check if an integer is prime.
func AsyncIsPrime(test int64) <-chan bool {
	result := make(chan bool)
	go func(i int64) {
		defer close(result)
		result <- IsPrime(i)
	}(test)
	return result
}

//IsPrime : Check if an integer is prime.
func IsPrime(test int64) bool {
	result := true
	if test <= 1 {
		result = false
	} else if test == 3 {
		result = true
	} else if test%2 == 0 {
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

//AsyncIsTwoSidedPrime : Checks if an integer is two sided prime
func AsyncIsTwoSidedPrime(test int64) <-chan bool {
	result := make(chan bool)
	go func(i int64) {
		defer close(result)
		result <- IsTwoSidedPrime(i)
	}(test)

	return result
}

//IsTwoSidedPrime : Checks if an integer is two sided prime
func IsTwoSidedPrime(test int64) bool {
	numberOfDigits := len(strconv.FormatInt(test, 10))
	result := IsPrime(test)
	for i := 1; i < numberOfDigits && result; i++ {
		divisor := int64(math.Pow10(i))
		firstPart := (test % divisor)
		secPart := test / divisor
		result1, result2 := AsyncIsPrime(firstPart), AsyncIsPrime(secPart)
		result = <-result1 && <-result2
	}
	return result
}
