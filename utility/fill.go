package utility

import "fmt"

func Fill(supWater, supMilk, supBeans, supCups *int) {
	var addWater, addMilk, addBeans, addCups int
	fmt.Println("Write how many ml of water you want to add:")
	fmt.Scan(&addWater)
	fmt.Println("Write how many ml of milk you want to add:")
	fmt.Scan(&addMilk)
	fmt.Println("Write how many grams of coffee beans you want to add:")
	fmt.Scan(&addBeans)
	fmt.Println("Write how many disposable coffee cups you want to add:")
	fmt.Scan(&addCups)
	*supWater += addWater
	*supMilk += addMilk
	*supBeans += addBeans
	*supCups += addCups
}
