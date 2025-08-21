package main

import (
	"fmt"
	"reflect"
)

func main() {
	determineType(10)
	determineType("smth")
	determineType(true)
	determineType(make(chan int))
}

func determineType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	default:
		if reflect.TypeOf(v).Kind() == reflect.Chan {
			fmt.Println("chan")
		}
	}
}
