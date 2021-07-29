package main

import (
	"fmt"
	"github.com/asishcse60/go-test-bible/calculator"
	"log"
	"net/http"
)

func main() {

	//go race.UpdateCounter()
	//fmt.Println("Here I am")
	Setup()
	fmt.Println(calculator.CalculateIsArmstrong(371))
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func Setup() {
	http.HandleFunc("/", HomePage)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome to homepage")
	fmt.Println("Homepage endpoints hits.")
}
