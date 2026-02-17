package user
import(

	"fmt"
	"time"
	// "errors"
)

type Admin struct {
	email string	
	password string
	User User
	
}

type User struct { // Here capital letter is used to make it public and accessible outside the package
	FirstName string // If we make a field start with small letter it will be private and not accessible outside the package
	// If it is capital letter it is public and accessible outside the package
	lastName  string
	birthdate string
	CreatedAt time.Time
}
func OutputUserDetails(u *User) {
	// (*u).firstName => Technically correct way of accessing a value stored in pointer by dereferencing it
	fmt.Print(u.FirstName)
	fmt.Println(u.lastName)
	fmt.Println(u.birthdate)
	// fmt.Println(time.Now())
}

func NewAdmin (email , password string) (Admin) {
	// if (fmt.Sprint(email) == "" || fmt.Sprint(password) == "") {
	// 	return nil, fmt.Errorf("Please provide all the details")
	// }
	return Admin{ // constructor is.used to return its address instead of value of the struct
		email: email,
		password: password,
		User : User{
			FirstName: "ADMIN",
			lastName: "ADMIN",
			birthdate: "---",
			CreatedAt: time.Now(),
		},

		
	}
}
func NewUser (userfirstName , userlastName ,userbirthdate string) (*User, error) {
    // Using constructr function not only return user it is also easy to validate the data and return error if any of the details is missing
	if (fmt.Sprint(userfirstName) == "" || fmt.Sprint(userlastName) == "" || fmt.Sprint(userbirthdate) == "") {
		return nil, fmt.Errorf("Please provide all the details")
	}
	return &User{ // constructor is.used to return its address instead of value of the struct
		FirstName: userfirstName,
		lastName:  userlastName,
		birthdate: userbirthdate,
		CreatedAt: time.Now(),
	},nil
}

// Now this is a method
func (u User) OutputDetails() {
	fmt.Printf("First name: %v\nLast name: %v\n DOB: %v\n",u.FirstName, u.lastName, u.birthdate)
}

func (u *User) ClearUserDetails(){ // pointer is used to clear user details 
	// now it will be added to original address to clear data
	u.FirstName=""
	u.lastName=""
}