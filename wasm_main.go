// +build js,wasm

package main

import (
	"fmt"
	"syscall/js"
)

func newTest(this js.Value, args []js.Value) interface{} {
	fmt.Println(args[0].Get("length"))

	// fmt.Println(args[0][1])
	for i := 0; i < args[0].Get("length").Int(); i++ {
		fmt.Println(args[0].Index(i))
	}

	c := []byte{1}
	this.Set("buf", js.TypedArrayOf(c))
	this.Set("change", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		c[0] = 9
		return nil
	}))
	return this
}

func main() {
	fmt.Println("hello!!!")
	window := js.Global()
	window.Set("Test", js.FuncOf(newTest))
	select {}
}
