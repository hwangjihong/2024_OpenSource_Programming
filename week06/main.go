package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 13
	var f float64 = 12.9
	c1 := 'A' // ASCII
	c2 := '김' // UNI

	fmt.Printf("value i : %d, value f : %f\n", i, f)
	// fmt.Printf("%d * %f = %f\n", i, f, i*f) invalid operation: i * f (mismatched types int and float64)
	fmt.Printf("%d * %f = %d\n", i, f, i*int(f)) // conversion
	// fmt.Printf("%d * %f = %f\n", i, f, float64(i)*f) // conversion
	fmt.Println(c1, c2)
	fmt.Println(reflect.TypeOf(i), reflect.TypeOf(f), reflect.TypeOf(c1), reflect.TypeOf(c2))

	var f1 float64
	var i1 int
	var b bool
	var s string

	fmt.Println(f1, b, s, i1)
	fmt.Printf("%f %t %s %d\n", f1, b, s, i1) // zero value
	f1 = 2.7
	i1 = 3
	fmt.Print(f1 > float64(i1), "\n") // comparsion (true/false)
}
