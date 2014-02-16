package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    if len(os.Args) == 1 {
        fmt.Println("Please supply an argument")
        os.Exit(1)
    }	
    fahrenheit, _ := strconv.ParseFloat(os.Args[1], 64)
    celsius := (fahrenheit - 32) * 5 / 9
    fmt.Printf("%g Fahrenheit = %g Celsius\n", fahrenheit, celsius)
}
