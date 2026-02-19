package main

import "fmt"

func generic(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generic("Hello, World!")
	generic(42)
	generic(3.14)
	generic(true)

}
