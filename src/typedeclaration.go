package main

import "fmt"

func main(){
	type Money uint

	var saldoJohn Money = 50000000
	var saldoDoe uint = 20000000

	var saldoDoeToMoney Money = Money(saldoDoe)
	
	fmt.Println(saldoJohn)
	fmt.Println(saldoDoeToMoney)
}