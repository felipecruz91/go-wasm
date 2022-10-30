//go:build js && wasm

package main

import "fmt"

func main() {
	fmt.Printf("Hi WebAssembly! 5+5=%d\n", 5+5)
}
