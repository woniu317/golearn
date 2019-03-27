package main

import (
	"fmt"
)

type ValueAndPointMethod struct {
	Va int
}

func (val *ValueAndPointMethod) pointSetVa(v int) {
	val.Va = v
}

func (val ValueAndPointMethod) valueSetVa(v int) {
	val.Va = v
}

func (val *ValueAndPointMethod) sayHello(string) string {
	fmt.Println("hello")
	return ""
}

func (val ValueAndPointMethod) sayWorld() {
	fmt.Println("world")
}

func main() {
	//val := ValueAndPointMethod{Va: 1}
	//val.sayWorld()
	//val.sayHello("hello")
	//fmt.Println("val's type:", reflect.TypeOf(val))
	vv := new(ValueAndPointMethod)
	//vv.sayHello("")
	//vv.sayWorld()
	//fmt.Println("vv's type:", reflect.TypeOf(vv))

	//fmt.Println("val's Va is ", val.Va)
	//val.pointSetVa(2)
	//fmt.Println("val's Va is ", val.Va)
	//val.valueSetVa(3)
	//fmt.Println("val's Va is ", val.Va)

	fmt.Println("vv's Va is ", vv.Va)
	vv.pointSetVa(2)
	fmt.Println("vv's Va is ", vv.Va)
	vv.valueSetVa(3)
	fmt.Println("vv's Va is ", vv.Va)

}
