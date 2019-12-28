# two-sided-prime
A project on Golang to determine if an integer is two sided prime.

### A number is a [Two-sided Prime](https://prime-numbers.info/article/two-sided-primes) if 
1.	The number is itself a prime number.  e.g. 3797
2.	Removing 1 digit at a time from left hand side from the original number gives you prime numbers. e.g. 797, 97, 7 are all prime numbers.
3.	Removing 1 digit at a time from right hand side from the original number gives you prime numbers. e.g. 379, 37, 3 are all prime numbers.

### Running the service
* In order to run program, first get the program using `go get github.com/radhe-soni/two-sided-prime/` and then
build program using 
`go build -o primeTestService github.com/radhe-soni/two-sided-prime/primetester/rest/`
* execute `./primeTestService` to start the service.
* Open (http://localhost:10000/health) in browser to check if service has started.
* Navigate to (http://localhost:10000/test/two-sided-prime/{number}) for utility end-point, replace "{number}" with test number.

