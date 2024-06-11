package main

import (
	"fmt"
	"reflect"
)

func main() {
	type Foo struct {
		A int     `myTag:"value"`
		B string  `myTag:"value2"`
		C float64 `myTag:-`
	}

	var f Foo
	ft := reflect.TypeOf(f)
	for i := 0; i < ft.NumField(); i++ {
		curField := ft.Field(i)
		fmt.Println(curField.Name, curField.Type.Name(), curField.Tag.Get("myTag"))
	}

	var i int8 = 10
	iv := reflect.ValueOf(&i)
	ivv := iv.Elem()
	fmt.Println(ivv.Kind())
	ivv.SetInt(500)
	fmt.Println(i)
	ivv.Set(reflect.ValueOf(int8(7)))
	fmt.Println(i)
}
