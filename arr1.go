package main

import (
	"fmt"
	"os"
	"sort"
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
	i := []int{a, b, c, d, e}
	sort.Ints(i)
	i[0], i[4] = i[4], i[0]
	fmt.Println(i)

}
