package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println("Hello World!")
	// var student string = "yamini"
	// fmt.Println(student)
	// fmt.Println("Hello")
	const inflation = 2.5
	var investment = 1000
	var returnval = 5.5
	var year = 10.0
	fmt.Print("Investment amount ")
	fmt.Scan(&investment)

	fmt.Print("Return value ")
	fmt.Scan(&returnval)

	fmt.Print("Year ")
	fmt.Scan(&year)

	future := float64(investment) * (math.Pow((1 + returnval/100), year))
	futurereal := future / math.Pow(1+(inflation/100), year)
	fmt.Printf("Future value %.0f \nFuture value(Adjusted for inflation): %.0f ", future, futurereal)
	// fmt.Println(future)
	// fmt.Println(futurereal)
}
