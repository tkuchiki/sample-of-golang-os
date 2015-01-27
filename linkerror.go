package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err := os.Link("/tmp", "/hogehogehoge")

	fmt.Printf("%T\n", err)
	fmt.Printf("%v\n", err.Error())

	if err != nil {
		log.Fatal(err)
	}
}
