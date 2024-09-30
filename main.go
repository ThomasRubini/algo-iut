//go:build !wasm

package main

import "algo-iut-1/internal/entrypoint"

func main() {
	entrypoint.Main()
}
