//Arrays! Generally not used as often compared to Slices, but slices are built on Arrays
//Different from C arrays in 3 ways:
//Arrays are values, assigning one array to another copies all elements
//passing arrays to functions will give the func a copy, not a pointer
//size of array is part of type. [10]int != [20]int

// but seriously, just use Slices
package main

// import "fmt"

// func main() {
// 	var a [2]string
// 	a[0] = "Hello"
// 	a[1] = "World"
// 	fmt.Println(a[0], a[1])
// 	fmt.Println(a)

// 	primes := [6]int{2, 3, 5, 7, 11, 13}
// 	fmt.Println(primes)
// }
