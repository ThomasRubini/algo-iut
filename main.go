//go:build !wasm

package main

import "algo-iut/internal/entrypoint"

func main() {
	entrypoint.Main()
}
