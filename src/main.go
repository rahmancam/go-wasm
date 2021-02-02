package main

import "syscall/js"

func add(this js.Value, inputs []js.Value) interface{} {
	return inputs[0].Float() + inputs[1].Float()
}

func subract(this js.Value, inputs []js.Value) interface{} {
	return inputs[0].Float() - inputs[1].Float()
}

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("subtract", js.FuncOf(subract))

}

func main() {
	c := make(chan int)
	println(("Go webAssembly Initialized"))
	registerCallbacks()

	<-c
}
