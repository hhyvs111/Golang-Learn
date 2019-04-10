package main

import "fmt"

func main() {

	i, j := 2, 2701

	p := &i
	//print the address
	fmt.Println(p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p
	p++
	fmt.Println(p)
	fmt.Println(j)
}
