package belajar_gorotines

import (
	"fmt"
	"testing"
	"time"
)

func runHelloWord() {
	fmt.Println("Hello Word")
}

func TestCreateGoroutines(t *testing.T) {
	go runHelloWord()
	fmt.Println("ups")

	time.Sleep(1 * time.Second)
}

func displayNumber(number int) {
	fmt.Println("ini adalah number", number)
}

func TestCreateGoroutinesMuch(t *testing.T) {
	for i := 0; i < 100; i++ {
		go displayNumber(i)
	}

	time.Sleep(10 * time.Second)
}
