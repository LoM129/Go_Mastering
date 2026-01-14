package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func handleSignal(signal os.Signal) {
	fmt.Println("handleSignal() Caught:", signal)
}

func main() {
	sigs := make(chan os.Signal, 1)
	//	signal.Notify(sigs, os.Interrupt, syscall.SIGINFO)
	signal.Notify(sigs, os.Interrupt)
	go func() {
		for {
			sig := <-sigs
			switch sig {
			case os.Interrupt:
				fmt.Println("Caught:", sig)
				//signal.Notify(sigs, os.Interrupt, syscall.SIGINFO)	case syscall.SIGINFO:
				handleSignal(sig)
				return
			}
		}
	}()

	//	for {
	fmt.Printf(".")
	time.Sleep(20 * time.Second)
	//	}
}
