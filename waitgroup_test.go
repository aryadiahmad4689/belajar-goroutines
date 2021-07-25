package belajar_gorotines

import (
	"fmt"
	"sync"
	"testing"
)

func TestWaitGroup(t *testing.T) {
	var group sync.WaitGroup

	for i := 0; i < 5; i++ {

		group.Add(1)
		go func() {
			fmt.Println("HALO INDONESIA")
		}()
		group.Done()

	}

	group.Wait()

}
