package main



import (
	"fmt"	
	"example.com/bank/fileops"
)
const accountBalancefile = "balance.txt"

func main() {
	for {
		var accountBalance, err = fileops.GetFloatfromfile(accountBalancefile)
		if err != nil {
			fmt.Println("Found errror")
			fmt.Println(err)
			fmt.Println("--------")
			//panic("Can't contnue sorry")
		}
		presentOptions()

		// var accountBalance = 1000.0

		var choice int
		fmt.Print("Your Choice ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("Your Balance is ", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Your Deposit Amount ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Your Balace has been updated! New amount: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance,accountBalancefile)
		case 3:

			var withdrawl float64
			fmt.Print("Amount of money you want to withdraw ")
			fmt.Scan(&withdrawl)
			if withdrawl <= 0 {
				fmt.Println("Invalid Amount")
				continue
			}
			if withdrawl > accountBalance {
				fmt.Println("You cant withdraw amount greater than you have")
				continue
			}
			accountBalance -= withdrawl
			fmt.Println("Your Balace has been updated! New amount: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance,accountBalancefile)
		default:
			fmt.Println("Bye ! Have a good day")
			return

		}

		// 	if choice == 1 {
		// 		fmt.Println("Your Balance is ", accountBalance)
		// 	} else if choice == 2 {
		// 		var depositAmount float64
		// 		fmt.Print("Your Deposit Amount ")
		// 		fmt.Scan(&depositAmount)
		// 		if depositAmount <= 0 {
		// 			fmt.Println("Invalid amount")
		// 			continue
		// 		}
		// 		accountBalance += depositAmount
		// 		fmt.Println("Your Balace has been updated! New amount: ", accountBalance)
		// 		writeBalance(accountBalance)

		// 	} else if choice == 3 {
		// 		var withdrawl float64
		// 		fmt.Print("Amount of money you want to withdraw ")
		// 		fmt.Scan(&withdrawl)
		// 		if withdrawl <= 0 {
		// 			fmt.Println("Invalid Amount")
		// 			continue
		// 		}
		// 		if withdrawl > accountBalance {
		// 			fmt.Println("You cant withdraw amount greater than you have")
		// 			continue
		// 		}
		// 		accountBalance -= withdrawl
		// 		fmt.Println("Your Balace has been updated! New amount: ", accountBalance)
		// 		writeBalance(accountBalance)

		// 	} else {
		// 		fmt.Println("Bye ! Have a good day")
		// 		break
		// 	}
		// }
		fmt.Println("Thanks for choosing our bank")

	}
}


