package main

import "fmt"

func main() {
	logger1 := GetInstance()
	logger2 := GetInstance()

	logger1.Log("This is a log message.")

	fmt.Println(logger1 == logger2) // true
}
