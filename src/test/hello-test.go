package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println(os.Getenv("os"))
	for _, env := range os.Environ() {
		fmt.Println(env)
	}

}
