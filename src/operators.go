package main

import "fmt"

func main(){
	a, b := 10, 3

    // Arithmetic Operators
    fmt.Println("Addition:", a+b)       // Output: 13
    fmt.Println("Subtraction:", a-b)    // Output: 7
    fmt.Println("Multiplication:", a*b) // Output: 30
    fmt.Println("Division:", a/b)       // Output: 3
    fmt.Println("Modulus:", a%b)        // Output: 1

    // Relational Operators
    fmt.Println("Equal to:", a == b)    // Output: false
    fmt.Println("Not Equal to:", a != b) // Output: true
    fmt.Println("Greater than:", a > b)  // Output: true
    fmt.Println("Less than:", a < b)     // Output: false
    fmt.Println("Greater than or equal:", a >= b) // Output: true
    fmt.Println("Less than or equal:", a <= b)    // Output: false

    // Logical Operators
    fmt.Println("Logical AND:", (a > b) && (b > 1))   // Output: true
    fmt.Println("Logical OR:", (a < b) || (b > 1))    // Output: true
    fmt.Println("Logical NOT:", !(a == b))             // Output: true

    // Bitwise Operators
    fmt.Println("Bitwise AND:", a&b)    // Output: 2 (00001010 & 00000011 = 00000010)
    fmt.Println("Bitwise OR:", a|b)     // Output: 11 (00001010 | 00000011 = 00001011)
    fmt.Println("Bitwise XOR:", a^b)    // Output: 9 (00001010 ^ 00000011 = 00001001)
    fmt.Println("Bitwise AND NOT:", a&^b) // Output: 8 (00001010 &^ 00000011 = 00001000)
    fmt.Println("Left Shift:", a<<1)    // Output: 20 (00001010 << 1 = 00010100)
    fmt.Println("Right Shift:", a>>1)   // Output: 5 (00001010 >> 1 = 00000101)
}