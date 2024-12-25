package main

import (
	"fmt"
	"reflect"
)

func main(){

	fmt.Println("one : ", 1, " (", reflect.TypeOf(1), ")")
	fmt.Println("two : ", 2,  " (", reflect.TypeOf(2), ")")
	fmt.Println("800 : ", " (", reflect.TypeOf(800), ")")
	fmt.Println("-2000 : ", " (", reflect.TypeOf(-2000), ")")

	fmt.Println("three point zaro : ", 3.0,  " (", reflect.TypeOf(3.0), ")")

	fmt.Println("true : ", " (", reflect.TypeOf(true), ")")
	fmt.Println("false : ", " (", reflect.TypeOf(false), ")")

	fmt.Println("string : ", " (", reflect.TypeOf("string : "), ")")
}