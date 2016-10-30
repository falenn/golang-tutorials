package main

//Programs start running in package main

// imports packages that are the name of the
// file with the same name as the last element
// in the path

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  //without seeding the random generator, you
  //will get the same random number.  Invoking rand
  //returns a random sequence, yet the result is
  //deterministic, such that the same sequence results without
  //"seeding" a different starting point for random generation.
  rand.Seed(time.Now().UTC().UnixNano())
  fmt.Println("May the force be with you")
  fmt.Println("My favorite number is ",
    rand.Intn(10))

  //fmt.Println and rand.Intn are "exported" funcs - they
  //are "public" from an OO point-of-view.

  //math.Pi is exported - it is a public constant.

  //math.pi is not exported and (well, doesn't exist) - but
  //if pi did exist in math, it would not be visible.

}
