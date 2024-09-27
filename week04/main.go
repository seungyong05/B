// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt" //c language #include <stdio.h>
	"math"
	"reflect"
	"strings"
)

func main() {
	//var i int
	//i = 55

	//var i int = 55

	var f float64 = 1.92

	//f := 1.92
	i := "55"
	fmt.Printf("%f %f\n", f, math.ceil(f))
	fmt.Println(strings.Title("inha"))
	fmt.Println(reflect.TypeOf(f), reflect.TypeOf(i))
	fmt.Println("Hello, world", i)
	fmt.Print("Hello, world", i, "\n")
	fmt.Println("Hello, world", i)
	fmt.Printf("Hello, world %s\n", i)
}
