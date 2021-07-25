package belajar_gorotines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChanel(t *testing.T) {
	chanel := make(chan string)

	// go func() {
	// 	chanel <- "halo ariadi ahmad"
	// 	fmt.Println("halo  makasar")
	// 	time.Sleep(2 * time.Second)
	// }()

	go giveMeResponse(chanel)

	data := <-chanel

	fmt.Println(data)

}

func giveMeResponse(chanel chan string) {

	chanel <- "halo indonesia"
	time.Sleep(50 * time.Second)
}

func TestChanelParameter(t *testing.T) {
	chanel := make(chan string)

	go giveMeResponse(chanel)

	data := <-chanel

	fmt.Println(data)

	close(chanel)
}

func onlyIn(chanel chan<- string) {
	chanel <- "ini saya berhaasil input data"
	time.Sleep(1 * time.Second)
}

func onlyOut(chanel <-chan string) {
	data := <-chanel
	fmt.Println(data)
}

func TestInOutChanel(t *testing.T) {
	chanel := make(chan string)

	go onlyIn(chanel)
	go onlyOut(chanel)

	time.Sleep(2 * time.Second)
	close(chanel)
}

func TestBufferedChanel(t *testing.T) {
	chanel := make(chan string, 3)

	go func() {
		chanel <- "halo indonesia"
		chanel <- "save palestine"
	}()

	go func() {
		fmt.Println(<-chanel)
		fmt.Println(<-chanel)
	}()

	time.Sleep(2 * time.Second)
	close(chanel)
}

func TestRangeChanel(t *testing.T) {
	chanel := make(chan string)
	go func() {
		for i := 0; i < 100; i++ {
			chanel <- strconv.Itoa(i)
		}
		defer close(chanel)
	}()

	for data := range chanel {
		fmt.Println(data)
	}
}

func TestSelectChanel(t *testing.T) {
	chanel1 := make(chan string)
	chanel2 := make(chan string)

	defer close(chanel1)
	defer close(chanel2)

	go giveMeResponse(chanel1)
	go giveMeResponse(chanel2)

	counter := 0
	for {
		select {
		case data := <-chanel1:
			fmt.Println("data dari chanel 1 ", data)
			counter++
		case data := <-chanel2:
			fmt.Println("data dari chanel 2 ", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}

}

func TestSelectChanelDefault(t *testing.T) {
	chanel1 := make(chan string)
	chanel2 := make(chan string)

	defer close(chanel1)
	defer close(chanel2)

	go giveMeResponse(chanel1)
	go giveMeResponse(chanel2)

	counter := 0
	for {
		select {
		case data := <-chanel1:
			fmt.Println("data dari chanel 1 ", data)
			counter++
		case data := <-chanel2:
			fmt.Println("data dari chanel 2 ", data)
			counter++
		default:
			fmt.Println("Sedang Menunggu Data")
		}

		if counter == 2 {
			break
		}
	}

}

func TestRaceCondition(t *testing.T) {
	var x = 0

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				x = x + 1
			}
		}()
	}

	time.Sleep(5 * time.Second)

	fmt.Println("jUMLAH", x)
}
