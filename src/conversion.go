package main

import "fmt"

func main(){
	var a int32 = 32768
	var b int64 = int64(a)
	var c int16 = int16(a) // number overflow

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	var fullName = "Muzammil nm"
	var getM = fullName[0]
	var mToString = string(getM)

	fmt.Println(fullName)
	fmt.Println(getM)
	fmt.Println(mToString)
}