package main

import (
	"fmt"
	"unsafe"
)

type emptyStruct struct {
}

func (em emptyStruct) sayHello() {
	fmt.Println("hello")
}

type emptyStruct2 struct {
}

func (em emptyStruct2) sayHello() {
	fmt.Println("hello2")
}

type oneParamStruct struct {
	id int
}

type twoParamsStruct struct {
	id   int
	dbid int
}

func (two twoParamsStruct) te() {

}

type threeParamsStruct struct {
	id   int
	dbid int
	Name string
}

func (th threeParamsStruct) String() string {
	fmt.Println("threeParamsStruct ...")
	return th.Name + "_" + string(th.dbid) + "_" + string(th.id)
}

func main() {
	//em := emptyStruct{}
	//em.sayHello()
	//fmt.Println(unsafe.Sizeof(em), reflect.TypeOf(em))
	//em2 := emptyStruct2{}
	//fmt.Println(unsafe.Sizeof(em2), reflect.TypeOf(em2))
	//em2.sayHello()

	a := int(1)
	fmt.Println(unsafe.Sizeof(a))
	ops := oneParamStruct{}
	fmt.Println(unsafe.Sizeof(ops))
	tps := twoParamsStruct{}
	fmt.Println(unsafe.Sizeof(tps))
	thps := threeParamsStruct{}
	fmt.Println(unsafe.Sizeof(thps))
	thps.Name = "hello word"
	fmt.Println(thps)
	str := "aaaaa"
	//fmt.Println(unsafe.Sizeof(str))
	fmt.Printf("%p\n", &str)
	emstr := "aaaaa"
	//fmt.Println(unsafe.Sizeof(emstr))
	fmt.Printf("%p\n", &emstr)
}
