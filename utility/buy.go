package utility

import (
	"fmt"
)

func SupplyCheck(state int, supWater, supMilk, supBeans, supCups *int) bool {
	var sufficient bool = true
	switch state {
	case 1:
		if *supWater-250 < 0 {
			fmt.Println("Sorry, not enough water!")
			sufficient = false
		} else if *supBeans-15 < 0 {
			fmt.Println("Sorry not enough coffee beans!")
			sufficient = false
		} else if *supCups-1 < 0 {
			fmt.Println("Sorry not enough disposable cups!")
			sufficient = false
		}
	case 2:
		if *supWater-350 < 0 {
			fmt.Println("Sorry, not enough water!")
			sufficient = false
		} else if *supMilk-75 < 0 {
			fmt.Println("Sorry not enough milk!")
			sufficient = false
		} else if *supBeans-20 < 0 {
			fmt.Println("Sorry not enough disposable cups!")
			sufficient = false
		} else if *supCups-1 < 0 {
			fmt.Println("Sorry not enough disposable cups!")
			sufficient = false
		}
	case 3:
		if *supWater-200 < 0 {
			fmt.Println("Sorry, not enough water!")
			sufficient = false
		} else if *supMilk-100 < 0 {
			fmt.Println("Sorry not enough milk!")
			sufficient = false
		} else if *supBeans-12 < 0 {
			fmt.Println("Sorry not enough disposable cups!")
			sufficient = false
		} else if *supCups-1 < 0 {
			fmt.Println("Sorry not enough disposable cups!")
			sufficient = false
		}
	}
	return (sufficient)
}

func Buy(supWater *int, supMilk *int, supBeans *int, supCups *int, remainMoney *int) bool {
	var state int
	var backToMain bool = false
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, 4 - to main menu:")
	fmt.Scan(&state)
	switch state {
	case 1:
		sufficientCheck := SupplyCheck(state, supWater, supMilk, supBeans, supCups)
		if sufficientCheck {
			fmt.Println("I have enough resources, making you a coffee!")
			*supWater -= 250
			*supBeans -= 16
			*remainMoney += 4
			*supCups -= 1
		}
	case 2:
		sufficientCheck := SupplyCheck(state, supWater, supMilk, supBeans, supCups)
		if sufficientCheck {
			fmt.Println("I have enough resources, making you a coffee!")
			*supWater -= 350
			*supMilk -= 75
			*supBeans -= 20
			*remainMoney += 7
			*supCups -= 1
		}
	case 3:
		sufficientCheck := SupplyCheck(state, supWater, supMilk, supBeans, supCups)
		if sufficientCheck {
			fmt.Println("I have enough resources, making you a coffee!")
			*supWater -= 200
			*supMilk -= 100
			*supBeans -= 12
			*remainMoney += 6
			*supCups -= 1
		}
	case 4:
		backToMain = true
	}
	return (backToMain)
}
