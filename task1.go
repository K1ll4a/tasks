// Задача 1
// ввод x и y -> координаты точки
// задание круга -> любое
// задание квадрата -> любое
// определить, находится ли точка внутри квадрата, внутри круга или и там, и там
// подсказка -> используй уравнение квадрата и уравнение окружности

// Задача 2
// Усложнение задачи 1
// Ты вводишь прямую
// Определить, находится ли прямая выше, ниже или внутри фигур
// Используй уравнение прямой
package main

import (
	"fmt"
	"os"
)

func main() {
	var x int
	fmt.Print("Enter the x coordinate: ")
	fmt.Fscan(os.Stdin, &x)
	var y int
	fmt.Print("Enter the y coordinate: ")
	fmt.Fscan(os.Stdin, &y)
	var y2 int
	fmt.Print("Enter the y coordinate of the straight line: ")
	fmt.Fscan(os.Stdin, &y2)
	var r int = 25
	var squareX1 int = 21
	var squareX2 int = 64
	var squareY1 int = 18
	var squareY2 int = 76

	if x*x+y*y <= r*r {
		fmt.Println("The point enters the circle")
	} else {
		fmt.Println("The point is not included in the circle")
	}
	if (x >= squareX1 && x <= squareX2) && (y >= squareY1 && y <= squareY2) {
		fmt.Println("The dot enters the square")
	} else {
		fmt.Println("The dot not enters the square")
	}
	if ((x >= squareX1 && x <= squareX2) && (y >= squareY1 && y <= squareY2)) && (x*x+y*y <= r*r) {
		fmt.Println("The point enters the square and the circle")
	}
	if y2 > squareY2 {
		fmt.Println("The line above the figures")
	}
	if y2 < -25 {
		fmt.Println("The line below the figures")
	}
	if y2 >= squareY1 && y2 <= 25 {
		fmt.Println("The line is included in both figures")
	}

}
