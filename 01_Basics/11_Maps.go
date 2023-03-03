package main

//maps aka dictionaries in C#, key value pairs.
//One of the 3 places you use Make() instead of New. See also Slices and Channels.
//In Go, maps can have any type of key so long as the keys can be compared for equality
//values can be any value
//keys and values must be of a single type each for a single map
//maps don't guarantee order of key-value pairs stored
//maps are passed by reference

// import "fmt"

// func main() {
// 	var timeZone = map[string]int{
// 		"UTC": 0 * 60 * 60,
// 		"EST": -5 * 60 * 60,
// 		"CST": -6 * 60 * 60,
// 		"MST": -7 * 60 * 60,
// 		"PST": -8 * 60 * 60,
// 	}
// 	fmt.Println(timeZone)

// 	//looking up a nonexistant key returns the zero value for given type.
// 	//problem: how to know if key nonexistant, or naturally returns 0?
// 	//answer: use "comma ok" idiom

// 	eastern, ok := timeZone["EST"]
// 	fmt.Println(eastern, ok) //ok is true
// 	western, ok := timeZone["WST"]
// 	fmt.Println(western, ok) //ok is false
// 	//does not HAVE to be "ok", can be any variable. Just a standard bool flag.
//	//can use _, ok if all you're looking for is if exists
// }
