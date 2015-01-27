package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var procAttr os.ProcAttr
	procAttr.Files = []*os.File{nil, os.Stdout, os.Stderr}

	proc, err := os.StartProcess("/bin/ls", []string{"-l"}, &procAttr)

	if err != nil {
		log.Fatal(err)
	}

	ps, err := proc.Wait()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Exited: %v \n", ps.Exited())
	fmt.Printf("Pid: %v \n", ps.Pid())
	fmt.Printf("String: %v \n", ps.String())
	fmt.Printf("Success: %v \n", ps.Success())
	fmt.Printf("Sys: %#v \n", ps.Sys())
	fmt.Printf("SystemTime: %v \n", ps.SystemTime())
	fmt.Printf("UserTime: %v \n", ps.UserTime())
}
