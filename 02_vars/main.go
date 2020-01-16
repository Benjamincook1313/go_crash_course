package main

import "fmt"


func main(){
	// MAIN TYPES
	// string
	// bool
	// int
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 unit64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	//  using var
	// var name = "Benjamin"
	var age int32 = 27
	var isCool = true
	isCool = false
	var size float32 = 2.3

	// shorthand
	// name := "Benjamin"
	// size := 1.3
	// email := "benjamin@gmail"

	name, email := "Ben", "Benjamin@gmail"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size)
}