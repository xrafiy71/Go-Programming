package main

import (
    "fmt"
)

// findMax finds the maximum number in a slice using mojo programming.
func findMax(numbers []int) int {
    if len(numbers) == 0 {
        return 0 // Default value for an empty slice
    }

    maxNum := numbers[0] // Assume the first number is the maximum

    for _, num := range numbers[1:] { // Iterate through the rest of the slice
        if num > maxNum { // If the current number is greater than the assumed maximum
            maxNum = num // Update the maximum number
        }
    }

    return maxNum // Return the maximum number
}

func main() {
    numbers := []int{5, 8, 2, 10, 3}
    fmt.Println("Maximum number in the slice:", findMax(numbers))
}
