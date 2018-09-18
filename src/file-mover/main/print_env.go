package main

import (
	"os"
	"fmt"
)

func main() {
	data := os.Environ()
	for _, val := range data {
		fmt.Println(val)
	}

	hostname, _ := os.Hostname()
	fmt.Println("hostname:", hostname)
}
