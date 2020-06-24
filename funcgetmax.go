/*Get_Max_Of_2_Numbers using function */
package main

import "fmt"

func getmax(x int, y int) int {
	var z int
	if x > y {
		z = x
	} else {
		z = y
	}
	return z
}
func main() {
	var a, b, max int
	fmt.Println("enter value of a ")
	fmt.Scanln(&a)
	fmt.Println("enter value of b ")
	fmt.Scanln(&b)
	max = getmax(a, b)
	fmt.Println("max value is", max)

}
