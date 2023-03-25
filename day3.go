package main
import (
    "fmt"
    "math"
    )

func main() {
    var mode int
    fmt.Println("Select an option \n 1.Rectangle \n 2.Square \n 3.Triangle")
    fmt.Scan(&mode)
    if mode == 1 {
        fmt.Println("-----AREA OF RECTANGLE-----")
        fmt.Println("Enter Length")
        var length float64
        fmt.Scan(&length)
        fmt.Println("Enter Breadth")
        var breadth float64
        fmt.Scan(&breadth)
        area := length * breadth
        fmt.Println("Area of Rectangle =", area)
    } else if mode == 2 {
        fmt.Println("-----AREA OF SQUARE-----")
        fmt.Println("Enter length of side")
        var side float64
        fmt.Scan(&side)
        area := side * side
        fmt.Println("Area of Square =", area)
    } else if mode == 3 {
        fmt.Println("-----AREA OF TRIANGLE-----")
        fmt.Println("Enter side 1")
        var side1 float64
        fmt.Scan(&side1)
        fmt.Println("Enter side 2")
        var side2 float64
        fmt.Scan(&side2)
        fmt.Println("Enter side 3")
        var side3 float64
        fmt.Scan(&side3)
        s:= (side1+side2+side3)/2
        fmt.Println(s*(s-side1)*(s-side2)*(s-side3))
        area:= math.Sqrt(s*(s-side1)*(s-side2)*(s-side3)) 
        fmt.Println("Area Of Triangle = ", area)
    } else {
        fmt.Println("Invalid Input")
    }
}