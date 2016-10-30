// Lesson - BasicTypes

package main

import (
  "fmt"
  "math/cmplx"
)

// basic types
/*
 bool
 string
 int int8 int16 int32 int64
 uint uint8 uint16 uint32 uint64 uintptr
 byte (uint8)
 rune (int32) - represents a unicode code point
 float32 float64
 complex64 complex128

 int uint uintptr are sized according to the architecture compiled on
 */

// Can define a collection of variables
var (
  ToBe bool = false  // assign boolean with value set
  MaxInt uint64 = 1<<64-1
  z complex128 = cmplx.Sqrt(-5 + 12i)

  i int
  fl float64
  b bool
  s string

  // Examples of type conversion
  // Type conversion is explicit like in Java
  num int = 42
  num64 float64 = float64(num)
  unum uint = uint(num64)

)

func main() {
  // format for print statement
  const f = "%T(%v)\n"

  // prints Type (value)
  // inputs of print format are vars for Type and value
  fmt.Printf(f, ToBe, ToBe)
  fmt.Printf(f, MaxInt, MaxInt)
  fmt.Printf(f, z, z)

  // These variables display zero values
  // zero values are what is set by default
  fmt.Printf(f, i, i)
  fmt.Printf(f, fl, fl)
  fmt.Printf(f, b, b)
  fmt.Printf(f, s, s)

  // print of num with type conversions
  fmt.Printf(f, num, num)
  fmt.Printf(f, num64, num64)
  fmt.Printf(f, unum, unum)

  // Variables can demonstrate type inference
  intVar := 33  // inferred type is int
  floatVar := 3.14159 // inferred float
  // print type inference
  fmt.Printf(f, intVar, intVar)
  fmt.Printf(f, floatVar, floatVar)
}
