package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}
type Manager struct {
	User
	title string
}

func (u User) Hello() {
	fmt.Println("Hello world")
}

func main() {
	u := User{1, "OK", 12}
	Info(&u)
	m := Manager{User: User{1, "OK", 13}, title: "123"}
}
func Info(o interface{}) {
	t := reflect.TypeOf(o)
	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("xx")
		return
	}
	v := reflect.ValueOf(o)
	fmt.Println("Fields:")
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()

		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}
}
