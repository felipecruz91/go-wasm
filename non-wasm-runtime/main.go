//go:build js && wasm

//go:generate cp $GOROOT/misc/wasm/wasm_exec.js .
package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	result := js.Global().Get("document").Call("querySelector", "#result")
	result.Set("innerHTML", fmt.Sprintf("Hi WebAssembly! 5+5=%d", 5+5))
}
