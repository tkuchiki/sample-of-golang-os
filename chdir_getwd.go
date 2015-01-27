package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err := os.Chdir("/tmp")

	if err != nil {
		log.Fatal(err)
	}

	dir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dir)
}
