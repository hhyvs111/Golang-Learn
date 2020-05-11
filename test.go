package main

import (
	"fmt"
	"time"
)
func main() {
	currentTime := time.Now().Format("20060102")
	fmt.Println(time.Now().Format("20060102"))
	fmt.Println(currentTime)
}
