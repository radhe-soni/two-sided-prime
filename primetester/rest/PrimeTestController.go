package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	primetest "github.com/radhe-soni/two-sided-prime/primetester/api"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Service is running!")
	fmt.Println("Endpoint Hit: homePage")
}
func isTwoSidedPrime(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["integer"]

	number, error := strconv.ParseInt(key, 10, 64)
	if error != nil {
		log.Fatal("Error occured while converting ", key, error)
	} else {
		fmt.Printf("%T %v", number, number)
		fmt.Fprintf(w, strconv.FormatBool(primetest.IsTwoSidedPrime(number)))
	}

}
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/health", homePage)
	myRouter.HandleFunc("/test/two-sided-prime/{integer}", isTwoSidedPrime)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	handleRequests()
}
