/*hello.go - My first Golang program */
package main

import (
	"fmt"
)

func main() {
	n, err := fmt.Printf("Hello, megha\n")
	fmt.Printf("%d %e", n, err)
	if err == nil {
		fmt.Printf("err is null")
	}

}
