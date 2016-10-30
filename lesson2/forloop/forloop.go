package main

import "fmt"

func main() {
  // standard loop syntax
  sum :=0
  fmt.Printf("count numbers from %v to %v\n", sum,10)
  for i :=0; i < 10; i++ {
    sum += 1
  }
  fmt.Println(sum)

  // minimized syntax on the for loop
  sum = 1
  fmt.Printf("Add numbers from %v to %v\n", sum,100)
  // for loop with a "while" feel
  for ; sum < 100; {
    sum += sum
  }
  fmt.Println(sum)

}
