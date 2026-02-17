package main

// If we use pointers we can avoid creating more memory
//Avoid unnecessary copies
//Directly mutate values
import "fmt"

func main() {
	age := 32
	var agePointer *int
	agePointer = &age
	fmt.Println(age)
	fmt.Println(age)
	fmt.Println(getAdultage(agePointer))
}
func getAdultage(age *int) int {
	//fmt.Println(agePointer-18)
	// We cant do arithmatic operations to pointers
	*age = *age - 18 //we can use * to dereference it from memory
	fmt.Println(age)
	return *age
}
