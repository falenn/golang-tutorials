// lesson - constants

package main

import "fmt"

// constant is declared with const keyword
// Here is a constant with package scope but without type.
// Constants cannot be declared with the := syntax but are still inferred.
const Pi = 3.14159

const (
  // create huge number by bit shifting a 1 bit left 100 places
  Big = 1 << 100
  //shift right, to end up with a small number - the number 2.
  Small = Big >> 99
)

func needInt(x int) int {
  return x*10+1
}

func needFloat(x float64) float64 {
  return x * 0.1
}


func main() {
  // format for print statement
  const f = "%T(%v)\n"

  //Here is a constant with function scope
  const World = "world"
  fmt.Println("Hello", World)
  fmt.Println("Happy",Pi,"Day")

  // demonstrate that type is inferred at assignmnet
  // event if not using := syntax
  fmt.Printf(f,Pi,Pi)

  const True = true
  fmt.Println("Go rules?", True)

  fmt.Println(needInt(Small))
  fmt.Println(needFloat(Small))
  fmt.Println(needFloat(Big))

  // this will overflow int as Big is too large.
  //fmt.Println(needInt(Big))
}
