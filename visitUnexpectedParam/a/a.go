package a

import (
	_ "unsafe"
)

//go:linkname say a.say
//go:nosplit
func say(name string) string {
	return "hello, " + name
}

//go:linkname say2 roobo.com/golearn/visitUnexpectedParam/b.Hi
func say2(name string) string {
	return "hi, " + name
}

func A() {

}
