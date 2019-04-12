package main

import "fmt"

type Phone interface {
	call()
}

type iPhone struct {
}

type xiaomi struct {
}

func (xs iPhone) call() {
	fmt.Println("I am iPhone xs , I can call you!")
}

func (mi8 xiaomi) call() {
	fmt.Println("I am xiaomi8, I can't call you!")
}
func main() {
	var phone Phone

	phone = new(iPhone)
	phone.call()

	phone = new(xiaomi)
	phone.call()
}

