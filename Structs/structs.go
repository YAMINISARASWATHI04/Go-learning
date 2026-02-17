package main

import (
	"fmt"
	// "time"
	"example.com/hello/user"
)



func main() {
	userfirstName := getUserData(" Please enter your firstname: ")
	userlastName := getUserData(" Please enter your Lastname: ")
	userbirthdate := getUserData(" Please enter your Date of birth(DD/MM/YYYY): ")

	var appuser *user.User
	appuser ,err := user.NewUser(userfirstName,userlastName,userbirthdate)
	if err != nil {
		fmt.Println(err)
		return
	}	

	admin := user.NewAdmin("test@example.com","password123");
	admin.User.OutputDetails() // We can access the details of user struct using admin struct because it is embedded in admin struct
	admin.User.ClearUserDetails()
	admin.User.OutputDetails() // Even after using this method data appears and not returns an empty string because
	fmt.Println(" ")

	user.OutputUserDetails(appuser)
	// For methods no need to pass arguments
	appuser.OutputDetails() // We can access methods using structs
	appuser.ClearUserDetails()// EVen after using this method data appears and not returns an empty string because
	// Already struct is stored and it is not cleared
	fmt.Println(" ")
	appuser.OutputDetails()
	// Generallly we use objects to which we assign values to use them In the same way we access them
	// using the objects created for the sake of structs

}

func getUserData(inputxt string) string {
	var userdata string
	fmt.Print(inputxt)
	fmt.Scanln(&userdata)
	return userdata
}
