package main

//bonus lesson not on the roadmap!
//quickly showing tags for things like filling out forms.

// import (
// 	"fmt"
// 	"reflect"
// )

// // tags are not recognized by themselves! This is just a way of annotating data.
// // Need something else to come in, recognize the tags, and enforce the rules.
// type Animal struct {
// 	Name   string `required max:"100"`
// 	Origin string
// }

// func main() {
// 	b := reflect.TypeOf(Animal{})
// 	field, _ := b.FieldByName("Name")
// 	fmt.Println(field.Tag)
// }
