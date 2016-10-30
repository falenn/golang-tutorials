package main

import "fmt"

func main() {
  sum := 1

  // for is go's while-loop!  Drop the semicolons
  // so, while sum is less than 1000, sum += sum.
  for sum < 1000 {
    sum += sum
  }
  fmt.Println(sum)

  //loop forever
  //for {
    
  //}
}
