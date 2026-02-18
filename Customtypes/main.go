package main

import "fmt"

//here type key word not only used to create strcts also we can create customized datatypes

type customString string

func (text customString) log(){

	fmt.Println(text)
}
func main (){
	var name customString ="Ishu"
	name.log()
	
}
// Here it might appear to be not so useful but further it will be very much useful to define customdatatypes