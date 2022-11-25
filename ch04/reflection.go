package main

import (
	"fmt"
	"reflect"
)

type Secret struct {
	Username string
	Password string
}

type Record struct {
	Field1 string
	Field2 float64
	Field3 Secret
}

func main() {
	a := Record{"Data String value", -12.123, Secret{"Ajay", "kakodia"}}
	r := reflect.ValueOf(a)
	fmt.Println("String Value:", r.String())

	aType := r.Type()
	fmt.Printf("a type: %s\n", aType)
	fmt.Printf("Number of fields %d in %s \n", r.NumField(), aType)

	for i := 0; i < r.NumField(); i++ {
		fmt.Printf("\t%s ", aType.Field(i).Name)
		fmt.Printf("\twith type: %s ", r.Field(i).Type())
		fmt.Printf("\tand value _ %v\n", r.Field(i).Interface())

		k := reflect.TypeOf(r.Field(i).Interface()).Kind()
		if k.String() == "struct" {
			fmt.Println(r.Field(i).Type())
		}

		if k == reflect.Struct {
			fmt.Println(r.Field(i).Type())
		}
	}
}
