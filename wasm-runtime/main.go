//go:build js && wasm

//go:generate cp /usr/local/tinygo/targets/wasm_exec.js .

package main

import "fmt"

func main() {
	fmt.Printf("Hi WebAssembly! 5+5=%d\n", 5+5)
}
