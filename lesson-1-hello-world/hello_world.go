package main

import "fmt"

func Hello(greeting string) string {
	return "Hello " + greeting
}
func main() {
	fmt.Println(Hello("my twin"))
}
