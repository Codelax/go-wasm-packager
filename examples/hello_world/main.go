package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World !")
	/*filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		fmt.Println(path)
		return nil
	})*/
}
