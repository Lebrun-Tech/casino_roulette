package main

import "fmt"

// Get the user Name
func GetName() string {
	name := ""
	
	fmt.Println()
	fmt.Println("Welcome to LEBRUN'S Casino...")
	fmt.Println("#############################")
	fmt.Println()
	fmt.Printf("Enter your name: ")


	fmt.Printf("%s", name)
	_, err := fmt.Scanln(&name)
	if err !=nil{
		return ""
	}
	fmt.Printf("Welcome %s, Let's PLAY! \n", name)
	return name
}

func GetBet(balance uint) uint {
	var bet uint
	for true{
		fmt.Printf("Enter your bet, or 0 to QUIT (Balance = %d): " ,balance)
		fmt.Scan(&bet)

		if bet > balance {
			fmt.Println("Bet can't be larger than balance")
		}else{
			break
		}
	}
	return bet
}
