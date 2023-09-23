// wasmモジュールにコンパイル
// GOOS=js GOARCH=wasm go build -o main.wasm main.go

package main

import (
        "fmt"
        "syscall/js"
)

func main() {
        c := make(chan struct{}, 0)

        js.Global().Set("helloWorld", js.FuncOf(helloWorld))

        <-c
}

func helloWorld(this js.Value, p []js.Value) interface{} {
        return fmt.Sprint("hello world")
}
