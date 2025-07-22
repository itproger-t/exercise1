package main

import (
	"fmt"
	"reflect"
)

func switchType(val interface{}) string {
	switch reflect.TypeOf(val).Kind() {
	case reflect.Int:
		return "int"
	case reflect.String:
		return "string"
	case reflect.Bool:
		return "bool"
	case reflect.Chan:
		return "chan"
	}
	return reflect.TypeOf(val).String()
}

func main() {
	fmt.Println(
		switchType(0), switchType(""),
		switchType(true), switchType(make(chan int)),
	)
}
