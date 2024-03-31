package main

import (
	"fmt"
	"playground/iteration"
	"reflect"
)

func main() {
	result := iteration.Repeat("a", 5)

	typeOf := reflect.TypeOf(result)

	fmt.Println(typeOf)

}
