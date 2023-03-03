// Resizable Array, used for nearly everything you'd use an array for
// One of the three places you use Make() for instead of New. Also in Maps and Channels.
package main

// import "fmt"

// func main() {
// 	primes := [6]int{2, 3, 5, 7, 11, 13} //array, not slice.
// 	j := []int{2}                        //slice.
// 	var s []int = primes[1:4]            //also slice. Note bounds are inclusive low, exclusive high.
// 	//can also use an index out of bounds for high to get final member.

// 	fmt.Println(s, j)

// 	j = append(j, 12) //adding to a slice. Can exceed underlying array's capacity. Copies into double size.
// 	fmt.Println(j)

// 	//Here's where it gets dangerous
// 	a := []int{1, 2, 3, 4, 5}
// 	fmt.Println(a)               //[1 2 3 4 5]
// 	b := append(a[:2], a[3:]...) //intention: b gets everything from a except middle value (3)
// 	fmt.Println(b)               //[1 2 4 5]
// 	fmt.Println(a)               //[1 2 4 5 5]
// 	//underlying arrays from slices get passed as references from one slice to another
// 	//my current guess is that the array started at 12345
// 	//then when you reference b, it only gives you the first 4 elements that have changed
// 	//but the final element is still there in memory, we haven't forced a reallocation
// 	//so the 5 is leftover from a, and the initial 1245 is from b
// 	//resulting in 1245 5

// 	c := make([]int, 10, 100) //makes slice of length 10, capacity 100
// 	fmt.Println(c)

// }
