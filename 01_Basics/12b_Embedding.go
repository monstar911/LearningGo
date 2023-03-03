package main

//bonus lesson not on the roadmap!
//Go is not object oriented; does not have inheritance
//handles that functionality differently. One limited way is through embedding.
//More graceful to handle through Interfaces, but this has some use still.

// import "fmt"

// type Animal struct {
// 	Name   string
// 	Origin string
// }

// //this is not an inheritance relationship - false to say bird IS AN animal
// //this is a composition relationship - bird HAS AN animal, or has animal characteristics
// //as mentioned, see Interfaces for more common solve for this problem.
// type Bird struct {
// 	Animal   //Embedding here
// 	SpeedKPH float32
// 	CanFly   bool
// }

// func main() {
// 	b := Bird{}
// 	b.Name = "Emu"
// 	b.Origin = "Australia"
// 	b.SpeedKPH = 48
// 	b.CanFly = false
// 	fmt.Println(b)
// }
