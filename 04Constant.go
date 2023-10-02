// Here's an example of how to declare and use constants in Go:

package main

import "fmt"

func main() {
    // Declaring a constant for pi
    const pi = 3.14159265359

    // Declaring multiple constants
    const (
        daysInWeek    = 7
        hoursInDay    = 24
        minutesInHour = 60
    )

    // Using constants in calculations
    circumference := 2 * pi * 5.0 // Calculating the circumference of a circle with radius 5

    // Printing the constants and the result
    fmt.Printf("Pi: %f\n", pi)
    fmt.Printf("Days in a week: %d\n", daysInWeek)
    fmt.Printf("Hours in a day: %d\n", hoursInDay)
    fmt.Printf("Minutes in an hour: %d\n", minutesInHour)
    fmt.Printf("Circumference of a circle with radius 5: %f\n", circumference)
}
