package utility

import "fmt"

func Take(remainMoney *int) {
	fmt.Printf("I gave you $%d\n", *remainMoney)
	*remainMoney = 0
}
