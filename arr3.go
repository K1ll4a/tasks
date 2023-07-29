// 3)
// Сдвинуть элементы массива на 1
// Ввод: ввод каждого значения массива/слайса
// Вывод: вывод результирующего массива/слайса
package main

import (
	"fmt"
	"math"
)

const N = 9

func main() {
	var arr [N]int
	var qty, i, j int
	fmt.Scanf("%d", &qty)
	for i = 0; i < N; i++ {
		arr[i] = (i+1)*100 + (i+1)*10 + (i + 1)
		fmt.Printf("%4d", arr[i])
	}
	fmt.Printf("\n")
	for j = 0; j < int(math.Abs(float64(qty))); j++ {
		if qty < 0 {
			for i = 0; i < N-1; i++ {
				arr[i] = arr[i+1]
			}
			arr[N-1] = 0
		} else {
			for i = N - 1; i > 0; i-- {
				arr[i] = arr[i-1]
			}
			arr[0] = 0
		}
		for i = 0; i < N; i++ {
			fmt.Printf("%4d", arr[i])
		}
		fmt.Printf("\n")
	}
}
