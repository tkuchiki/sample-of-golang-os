package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		os.Interrupt,
		os.Kill,
	)
	defer signal.Stop(sigChan)

	exitChan := make(chan int)
	go func() {
		for {
			s := <-sigChan
			switch s {
			case syscall.SIGHUP:
				fmt.Println("SIGHUP")
			case syscall.SIGINT:
				fmt.Println("SIGINT")
				exitChan <- 0
			case syscall.SIGTERM:
				fmt.Println("SIGTERM")
				exitChan <- 0
			case syscall.SIGQUIT:
				fmt.Println("SIGQUIT")
				exitChan <- 0
			default:
				fmt.Println("Unknown signal")
				exitChan <- 1
			}
		}
	}()

	os.Exit(<-exitChan)
}
