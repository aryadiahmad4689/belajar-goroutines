package belajar_gorotines

import (
	"fmt"
	"sync"
	"testing"
)

func tulis(data []string, ch chan string) {
	for _, dt := range data {
		ch <- dt
	}
	close(ch)
}

func cetak(ch chan string) {
	for result := range ch {
		fmt.Println(result)
	}
}

func TestChanel(t *testing.T) {
	var group sync.WaitGroup

	data := []string{"PHP", "Golang", "Javascript", "Python"}
	ch := make(chan string, len(data))
	group.Add(1)
	go tulis(data, ch)
	go cetak(ch)
	group.Done()

	group.Wait()
}

func sumData(data []int, ch chan int) {
	go func() {
		var nilai int

		for _, value := range data {
			nilai = nilai + value
		}
		ch <- nilai
		close(ch)

	}()

}

func TestSumData(t *testing.T) {
	var group sync.WaitGroup
	ch := make(chan int)
	data := []int{
		1, 2, 3, 4, 5, 6, 7,
	}
	group.Add(1)
	go sumData(data, ch)
	nilai := <-ch
	fmt.Println(nilai)
	fmt.Println()
	group.Done()

	group.Wait()

}
