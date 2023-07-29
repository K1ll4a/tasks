// Задача 3
// Дана общая стоимость товара.
// Написать программу вычисления суммы к оплате с учетом скидки.
// Скидка в 10% от стоимости товара предоставляется если об

package main

import (
	"fmt"
	"os"
)

func main() {
	var cost int
	fmt.Print("Enter the cost of the product: ")
	fmt.Fscan(os.Stdin, &cost)

	if cost > 500 {
		fmt.Println("Discount = ", cost*10/100, "Sum with discount = ", cost-cost*10/100)
	} else {
		fmt.Println("Sorry, you are not getting a discount")
	}
}
