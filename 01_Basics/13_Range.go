package main

//Range helps iterate over data structures

// import "fmt"

// func main() {
// 	//when using range to iterate over a slice, two values will be returned in each iteration.
// 	//Returns index and copy of value. Here, we don't care about index, so we underscore it _.
// 	nums := []int{2, 3, 4}
// 	sum := 0
// 	for _, num := range nums {
// 		sum += num
// 	}
// 	fmt.Println("sum:", sum)

// 	for i, num := range nums {
// 		if num == 3 {
// 			fmt.Println("index:", i)
// 		}
// 	}

// 	//here, range returns key and value when applied to maps.
// 	kvs := map[string]string{"a": "apple", "b": "banana"}
// 	for k, v := range kvs {
// 		fmt.Printf("%s -> %s\n", k, v)
// 	}

// 	//just keys here
// 	for k := range kvs {
// 		fmt.Println("key:", k)
// 	}

// 	//iterating over strings gets into unicode and runes, beyond my current understanding.
// 	for i, c := range "go" {
// 		fmt.Println(i, c)
// 		//yields 0 103, then 1 111. starting byte index on rune then rune itself.
// 		//whatever that means.
// 	}
// }
