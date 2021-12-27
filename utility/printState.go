package utility

import "fmt"

func PrintState(supWater, supMilk, supBeans, supCups, remainMoney *int) {
	fmt.Println("The coffee machine has:")
	fmt.Printf("%d of water\n", *supWater)
	fmt.Printf("%d of milk\n", *supMilk)
	fmt.Printf("%d of coffee beans\n", *supBeans)
	fmt.Printf("%d of disposable cups\n", *supCups)
	fmt.Printf("%d of money\n", *remainMoney)
}
