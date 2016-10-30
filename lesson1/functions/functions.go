// Lesson3 - Functions

package main

import "fmt"

// Create a function that adds two numbers
// together.
// The signature of the function is:
// add(x int, y int) int
// name:  add
// parameters: (x int, y int) - notice typing after variable name
// returns: int
func add(x int, y int) int {
  return x + y
}

// another function -
// When 2 or more parameters share the same type, you can
// specify the type with a shorthand notation
func subtract(x,y int) int {
  return x - y
}

// A function can return multiple items
// here, we will swap the order or x and y
func swap(x,y string) (string, string) {
  return y, x
}

// a function with a "naked return"  Here, there is no
// explicit returned variable - the value as set is returned.
// This is another shorthand, good in small functions  - bad for
// readability in larger functions.
func divideBy2(half int) (x int) {
  x = half/2
  return
}

// create main function
func main() {
  //call the add function and print the results.
  fmt.Println("add 42 and 13: ", add(42,13))
  fmt.Println("Subtract 42 from 13: ",subtract(42,13))
  fmt.Println("Swapping black and white")

  // assigning a, b to the output of swap
  a,b := swap("black", "white")
  fmt.Println("swapped order: ",a,b)
  fmt.Println("Divide 54 by 2: ",divideBy2(54))
}
