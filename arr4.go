// 4)
// Разделение отрицательных и положительных чисел с сохранением порядка
// Ввод: ввод каждого значения массива/слайса
// Вывод: 2 массива/слайся (с положительными и отрицательными значениями)
package main

import (
	"fmt"
	"os"
)

func main() {
	var a int
	fmt.Print("Enter first element: ")
	fmt.Fscan(os.Stdin, &a)
	var b int
	fmt.Print("Enter second element: ")
	fmt.Fscan(os.Stdin, &b)
	var c int
	fmt.Print("Enter third element: ")
	fmt.Fscan(os.Stdin, &c)
	var d int
	fmt.Print("Enter fourth element: ")
	fmt.Fscan(os.Stdin, &d)
	var e int
	fmt.Print("Enter fifth element: ")
	fmt.Fscan(os.Stdin, &e)
	arr := []int{a, b, c, d, e}
	arr1 := []int{}
	arr2 := []int{}

	for i := range arr {
		if arr[i] < 0 {
			arr1 = append(arr1, arr[i])

		}
		if arr[i] > 0 {
			arr2 = append(arr2, arr[i])
		}
	}
	fmt.Println("Array with negative numbers: ", arr1)
	fmt.Println("Array with positive numbers: ", arr2)
}
