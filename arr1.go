// 1)
// Необходимо поменять местами максимальный и первый элементы массива/слайса
// Ввод: ввод каждого значения массива/слайса
// Вывод: вывод результирующего массива/слайса
// Golang Program to Find Largest Element in an Array
package main

import (
	"fmt"
	"os"
)

func main() {
	var total, i, largest, pos int

	fmt.Print("Enter the number of array elements: ")
	fmt.Fscan(os.Stdin, &total)
	n := make([]int, total)

	fmt.Scanln(&total)
	for i = 0; i < total; i++ {
		fmt.Print("Enter the number : ")
		fmt.Scan(&n[i])

	}
	largest = n[0]
	for i = 0; i < total; i++ {
		if largest < n[i] {
			largest = n[i]
			pos = i

		}
	}
	n[0], n[pos] = n[pos], n[0]
	fmt.Println(n)

}
