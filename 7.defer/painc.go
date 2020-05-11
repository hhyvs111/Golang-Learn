package main

import "fmt"
func caller() {
	defer func() {
		fmt.Println("defered annoymous function in caller")
		if r := recover(); r != nil {
			fmt.Println("recovered from called")
		}
	}()
	callee()
}
func callee() {
	panic(111)
}
func main() {
	caller()
	fmt.Println("vim-go")
}


