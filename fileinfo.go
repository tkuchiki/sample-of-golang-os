package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fp, err := ioutil.TempFile("", "")

	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(fp.Name())
	defer fp.Close()

	fi, err := fp.Stat()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Name: %v \n", fi.Name())
	fmt.Printf("Size: %v \n", fi.Size())
	fmt.Printf("Mode: %v \n", fi.Mode())
	fmt.Printf("ModTime: %v \n", fi.ModTime())
	fmt.Printf("IsDir: %v \n", fi.IsDir())
	fmt.Printf("Sys: %#v \n", fi.Sys())
}
