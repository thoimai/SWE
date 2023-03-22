package main

import "fmt"

func main() {
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println("The sum of 100 consecutive numbers from 1 to 100 is:", sum)
}
