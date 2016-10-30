// Lesson 4 - variables
package main

import "fmt"

//keyword var declase a LIST of variables
//The type of these variables is the last item in the list:
// For instance, these are all of type Boolean (bool)
// These variables are scoped at the Package level.
var c, python, java bool

// Lets initialize a package variable
var i, j int = 1,2

func main() {
  //Let's declase a locally scoped, function-level
  // variable, i of type int.
  var i = 33

  // when we run this, we see that i is == 33!  This is the
  // value of i in the function scope.  This trumps package.
  // Also, we omitted the type when initializing i to 33 - that's
  // because the type is inferred.
  fmt.Println(i,j,c,python,java)
}
