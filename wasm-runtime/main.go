//go:build js && wasm

package main

import "log"

func main() {
	log.Printf("Hi WebAssembly! 5+5=%d", 5+5)
}
