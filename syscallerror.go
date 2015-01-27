package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	err := syscall.Chdir("/hogehogehoge")

	fmt.Printf("%T\n", err)

	syscallErr := os.NewSyscallError("chdir", err)

	fmt.Printf("%T\n", syscallErr)
	fmt.Printf("%v\n", syscallErr.Error())
}
