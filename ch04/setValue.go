// Changing value of variable in struct using refection
package main

import (
	"fmt"
	"reflect"
)

type T struct {
	F1 int
	F2 string
	F3 float64
}

func main() {
	a := T{1, "F2", 3.0}
	fmt.Println("A: ", a)
	// with the use of Elem() and Pointer to variable a, variable can be modified if needed
	r := reflect.ValueOf(&a).Elem()
	fmt.Println("String Value:", r.String())
	typeofa := r.Type()
	fmt.Println("Type of a: ", typeofa)
	for i := 0; i < r.NumField(); i++ {
		f := r.Field(i)
		tofA := typeofa.Field(i).Name
		fmt.Printf("%d: %s %s = %v\n", i, tofA, f.Type(), f.Interface())

		k := reflect.TypeOf(f.Interface()).Kind()
		if k == reflect.Int {
			// setting int value to -100
			r.Field(i).SetInt(-100)
		} else if k == reflect.String {
			// setting string value to Changed
			r.Field(i).SetString("Changed!!")
		}
	}
	fmt.Println("New a values: ", a)
}
