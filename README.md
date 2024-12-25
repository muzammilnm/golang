# Golang Beginner

## Introduction
Golang, also known as Go, is an open-source programming language developed by Google. Designed for efficiency, simplicity, and high performance, Go is widely used to build backend applications, cloud services, and distributed systems. With its robust concurrency support through goroutines, Go empowers developers to manage multiple tasks simultaneously with optimal performance. The language also boasts a rich ecosystem of libraries and built-in tools like gofmt to ensure code consistency.

### 1. Hello World
to display something in the terminal, we can use the `fmt` library

```go
package main

import ("fmt")

func main() {
  fmt.Println("Hello World!")
}
```

#### Explanation
`package main`: the indicates it as the main application
`func main()`: it is the main function that will be executed first when the application runs

### 2. Data Types
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

### 3. 
