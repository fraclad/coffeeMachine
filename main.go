package main

import (
	"coffeeMachine/utility"
	"fmt"
)

func main() {
	water := 400
	milk := 540
	beans := 120
	cups := 9
	money := 550

	var action string
	for action != "exit" {
		fmt.Println("\nWrite action (buy, fill, take, remaining, exit):")
		fmt.Scan(&action)
		switch action {
		case "remaining":
			utility.PrintState(&water, &milk, &beans, &cups, &money)
		case "buy":
			backMain := utility.Buy(&water, &milk, &beans, &cups, &money)
			//fmt.Println(backMain)
			if backMain {
				break
			}
		case "fill":
			utility.Fill(&water, &milk, &beans, &cups)
		case "take":
			utility.Take(&money)
		case "exit":
			fmt.Println("Thank you for using the coffee machine! ☕️")
		default:
			fmt.Println("Please enter a valid option!")
		}
	}
}
