package main

import "fmt"

func main() {
    // Declare and initialize an array
    numbers := [5]int{1, 2, 3, 4, 5}
	var name [2]string
	var country = [...]string{
		"Indonesia",
		"Japan",
		"Australia",
	}

    // Access an array element
    fmt.Println("First element:", numbers[0])

    // Modify an element
    numbers[4] = 10
	name[0] = "Muzammil"
	name[1] = "nm"
    fmt.Println("Updated array:", numbers)
	fmt.Println(name)
	fmt.Println(country)

    // Get the length of the array
    fmt.Println("Array length:", len(numbers))

    // Iterate through the array
    fmt.Println("Array elements:")
    for _, value := range numbers {
        fmt.Println(value)
    }
}
