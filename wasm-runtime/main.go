//go:build js && wasm

//go:generate cp /usr/local/tinygo/targets/wasm_exec.js .

package main

func main() {
	println("Hi WebAssembly! 5+5=%d", 5+5)
}
