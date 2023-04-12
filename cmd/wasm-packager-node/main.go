package main

import (
	"bytes"
	"flag"
	"log"
	"os"
	"wasm-packager/internal/wasmexec"
)

func main() {
	target := "standalone.js"
	wasmFile := "main.wasm"
	flag.StringVar(&target, "target", "standalone.js", "The target js file")
	flag.StringVar(&wasmFile, "wasm", "main.wasm", "The wasm binary to bundle")
	flag.Parse()

	fileContent := bytes.NewBuffer(nil)
	fileContent.WriteString(wasmexec.NodeScript())
	fileContent.WriteString(wasmexec.Script())

	f, err := os.Open(wasmFile)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	wasmBuffer, err := wasmexec.Buffer(f)
	if err != nil {
		log.Fatalln(err)
	}

	fileContent.WriteString(wasmBuffer)
	fileContent.WriteString(wasmexec.Runner())
	err = os.WriteFile(target, fileContent.Bytes(), os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
}
