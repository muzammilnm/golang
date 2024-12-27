# Golang Beginner

## Introduction
Golang, also known as Go, is an open-source programming language developed by Google. Designed for efficiency, simplicity, and high performance, Go is widely used to build backend applications, cloud services, and distributed systems. With its robust concurrency support through goroutines, Go empowers developers to manage multiple tasks simultaneously with optimal performance. The language also boasts a rich ecosystem of libraries and built-in tools like gofmt to ensure code consistency.

## 1. Hello World
to display something in the terminal, we can use the `fmt` library

```go
package main

import ("fmt")

func main() {
  fmt.Println("Hello World!")
}
```

#### Explanation
- `package main`: the indicates it as the main application
- `func main()`: it is the main function that will be executed first when the application runs

## 2. Data Types
Golang provides a rich set of data types to handle various kinds of data efficiently. understanding these typesis essential for writing robust and performant programs

#### Numeric Types (Integer)
- `int8` from -128 to 127
- `int16` from -32768 to 32767
- `int32` from -2147483648
- `int64` from -9223372036854775808 to 9223372036854775807
- `uint8` from 0 to 255
- `uint16` from 0 to 65535
- `uint32` from 0 to 4294967295
- `uint64` from 0 to 18446744073709551615

#### Numeric Types (float)
- `float32` from -3.4e+38 to 3.4e+38
- `float64` from -1.7e+308 to +1.7e+308
- `complex64` complex numbers with `float32` real and imaginary parts
- `complex128` complex numbers with `float64` ral and imaginary parts

#### Numeric Alias
- `byte` alias for `uint8`
- `rune` alias for `int32`
- `int` alias for minimum `int32`
- `uint` alias for minimum `uint32`

#### Boolean
- `true`
- `false`

#### String
immutable sequences of bytes that represent text

```go
  fmt.println("iam learn golang")
```

## 3. Variable
variable are fundamental in any programming language, including Golang. they are used to store and manage data during the execution of a program. in Go, variable have specific characteristics and can be declared in multiple ways to suit various use cases.

```go
var name string
name = "Muzammil"

var address string = "Jakarta"
var gender = "male"
```

besides using `var`, we can also use the `:=` syntax when creating variables

```go
  email := "email@email.com"
```

you can declare and initialize multiple variables in a single line

```go
var x, y, z int = 1, 2, 3
var (
  firstName = "Muzammil"
  lastName = "nm"
)
a, b := "Go", "Language"
```

## 4. Constant
constants are immutable values that are determined at complile-time and cannot be changed during the excetion of the program.

```go
const firstName = "Muzammil"
const lastName string = "nm"

lastName = "nm" // cannot assign to lastName
```

## 5. Type Conversion
type convertion is the process  of explicitly converting a value from one data type to another

```go
var a int32 = 32768
var b int64 = int64(a)
var c int16 = int16(a) // number overflow
```

if you want to convert a number to a string, you can do it like this

```go
var fullName = "Muzammil nm"
var getM = fullName[0]
var mToString = string(getM)
```

## 5. Type Declarations
type declaration allows developers to crate custom types that are derived from existing types. this feature is powerful for improving code readability, creating domain-specific types, and iplementing methods on custom types.

```go
type Money uint

var saldoJohn Money = 50000000
var saldoDoe uint = 20000000

var saldoDoeToMoney Money = Money(saldoDoe)
```

## 6. Operators
operators are symbols or keyword that perform operation on operands. they are used to manipulate data, compare values, and control the flow of execution

### Arithmetic Operators
Used for mathematical calculations.
| Operator | Description	       | Example   |
| :------- | :-----------------: | :------   |
| `+`      | Addition 	         | `a + b`   |
| `-`      | Subtraction	       | `a - b`   |
| `*`      | Multiplication      |	`a * b`  |
| `/`	     | Division	           | `a / b`   |
| `%`      | Modulus (Remainder) |	`a % b`  |

Example :
```go
a, b := 10, 3

fmt.Println("Addition:", a+b)       // Output: 13
fmt.Println("Subtraction:", a-b)    // Output: 7
fmt.Println("Multiplication:", a*b) // Output: 30
fmt.Println("Division:", a/b)       // Output: 3
fmt.Println("Modulus:", a%b)        // Output: 1
```

### Relational Operator
used to compare two values, returning a boolean (`true` or `false`)

| Operator | Description	            | Example  |
| :------- | :---------------------:  | :------  |
| `==`     | Equal to	                | `a == b` |
| `!=`     | Not equal to	            | `a != b` |
| `>`	     | Greater than	            | `a > b`  |
| `<`	     | Less than	              | `a < b`  |
| `>=`	   | Greater than or equal to |	`a >= b` |
| `<=`	   | Less than or equal to    | `a <= b` |

Example :
```go
a, b := 10, 3

fmt.Println("Equal to:", a == b)    // Output: false
fmt.Println("Not Equal to:", a != b) // Output: true
fmt.Println("Greater than:", a > b)  // Output: true
fmt.Println("Less than:", a < b)     // Output: false
fmt.Println("Greater than or equal:", a >= b) // Output: true
fmt.Println("Less than or equal:", a <= b)    // Output: false
```

### Logical Operator
used to perform logical operations, commonly in conditional statements

|Operator |	Description	|Example     |
| :-----  | :-------:   | :------    |
| `&&`	  | Logical AND	| `a && b`   |
| `\|\|`  | Logical OR  | `a \|\| b` |
| `!`	    | Logical NOT	| `!a`       |

Example :
```go
a, b := 10, 3

fmt.Println("Logical AND:", (a > b) && (b > 1))   // Output: true
fmt.Println("Logical OR:", (a < b) || (b > 1))    // Output: true
fmt.Println("Logical NOT:", !(a == b))             // Output: true
```

### Assignment Operator
used to assign values to variables

| Operator |	 Description       |	Example                             |
| :------- | :----------------:  | :----------------------------------- |
|   `=`    | Simple assignment   | `a = b`                              |
|   `+=`   | Add and assign	     | `a += b` (equivalent to `a = a + b`) |
|   `-=`   | Subtract and assign | `a -= b` (equivalent to `a = a - b`) |
|   `*=`   | Multiply and assign | `a *= b` (equivalent to `a = a * b`) |
|   `/=`   | Divide and assign	 | `a /= b` (equivalent to `a = a / b`) |
|   `%=`   | Modulus and assign	 | `a %= b` (equivalent to `a = a % b`) |


### Bitwise Operator
perform operations on binary representations of integers

| Operator | Description          | Example                           | Result (Binary)      | Result (Decimal) |
|----------|----------------------|-----------------------------------|----------------------|------------------|
| `&`      | Bitwise AND          | `a & b` (5 & 3)                  | `00000001`           | `1`              |
| `|`      | Bitwise OR           | `a | b` (5 | 3)                  | `00000111`           | `7`              |
| `^`      | Bitwise XOR          | `a ^ b` (5 ^ 3)                  | `00000110`           | `6`              |
| `&^`     | Bit Clear (AND NOT)  | `a &^ b` (5 &^ 3)                | `00000100`           | `4`              |
| `<<`     | Left shift           | `a << 1` (5 << 1)                | `00001010`           | `10`             |
| `>>`     | Right shift          | `a >> 1` (5 >> 1)                | `00000010`           | `2`              |

Example :
```go
a, b := 10, 3

fmt.Println("Bitwise AND:", a&b)    // Output: 2 (00001010 & 00000011 = 00000010)
fmt.Println("Bitwise OR:", a|b)     // Output: 11 (00001010 | 00000011 = 00001011)
fmt.Println("Bitwise XOR:", a^b)    // Output: 9 (00001010 ^ 00000011 = 00001001)
fmt.Println("Bitwise AND NOT:", a&^b) // Output: 8 (00001010 &^ 00000011 = 00001000)
fmt.Println("Left Shift:", a<<1)    // Output: 20 (00001010 << 1 = 00010100)
fmt.Println("Right Shift:", a>>1)   // Output: 5 (00001010 >> 1 = 00000101)
```

## 7. Array
array is a data structure used to store a fixed number of elements of the same type in a contiguous block memory

### Declaring an Array
you can declare an array by specifying its size and type
```go
var arr [5]int
```

### initializing an Array
you can directly assign values during declaration
```go
var arr = [3]int{1,2,3}
```

### Using Ellipsis (`...`)
if the size is not specified, Golang can infer it using an ellipsis
```go
arr := [...]int{4,5,6}
```

### Accesing Array
Element in an array are accessed using their index, starting from 0
```go
arr = [3]string{"iam", "learn", "golang"}
fmt.Println(arr[0])
```

you can also modify elements by their index
```go
arr[0] = "you"
fmt.Println(arr)
```

### Getting the length of an array
use the `len()` function to determine the length of an array
```go
arr := [4]int{10, 20, 30, 40}
fmt.Println(len(arr)) // Output: 4
```

### Iterating over an array
Use a `for` loop to iterate through the elements
```go
arr := [3]int{10, 20, 30}
for i := 0; i < len(arr); i++ {
    fmt.Println(arr[i])
}
```

alternatively, use `for range` loop for a simpler approach
```go
for index, value := range arr {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}
```

## 8. Slice
slice in golang is lightweight and flexible abstraction of an array. while an array has a fixed size, a slice provides a way to work with sequences of elements that can grow or shrink dinamicaly

### creating and initializing slices
you can create a slice using the `[]` notation without specifying a size
```go
var slice []int
fmt.Println(slice) // Output: []
```

using `make()` function is to create a slice with a specific length and capacity
```go
slice := make([]int, 5, 10)
fmt.Println(slice) // Output: [0 0 0 0 0]
```

a slice can be created from an existing array
```go
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:4]
fmt.Println(slice) // Output: [2 3 4]
```

you can initialize a slice with values directly
```go
slice := []string{"Go", "is", "fun"}
fmt.Println(slice) // Output: [Go is fun]
```

### Accessing and Modifying slice elements
elements in a slice are accessed and modified using their index, similar to arrays
```go
slice := []int{10, 20, 30}
fmt.Println(slice[1]) // Output: 20
slice[1] = 25
fmt.Println(slice)    // Output: [10 25 30]
```

### Appending elements to a slice
you can use the `append()` function to add elements to a slice dinamically
```go
slice := []int{1, 2, 3}
slice = append(slice, 4, 5)
fmt.Println(slice) // Output: [1 2 3 4 5]
```

### Copy a slice
the `copy()` function is used to copy elements from one slice to another
```go
src := []int{1, 2, 3}
dest := make([]int, len(src))
copy(dest, src)
fmt.Println(dest) // Output: [1 2 3]
```

### Length and capacity
<b>length</b> is the number of elements in the slice. <b>capacity</br> is the total number of elements the slice can hold before reallocating memory
```go
slice := make([]int, 3, 5)
fmt.Println(len(slice)) // Output: 3
fmt.Println(cap(slice)) // Output: 5
```

### Slicing in slice
slices can be sliced further to create new slices
```go
slice := []int{1, 2, 3, 4, 5}
subSlice := slice[1:4]
fmt.Println(subSlice) // Output: [2 3 4]
```