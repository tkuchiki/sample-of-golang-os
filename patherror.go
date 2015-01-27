package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err := os.Chdir("/hogehogehoge")

	fmt.Printf("%T\n", err)
	fmt.Printf("%v\n", err.Error())

	if err != nil {
		log.Fatal(err)
	}
}
