package belajar_gorotines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestFixRaceCondition(t *testing.T) {
	var x = 0
	var mutex sync.Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)

	fmt.Println("jUMLAH", x)
}

type Account struct {
	RWMutex sync.RWMutex
	balance int
}

func (account *Account) addBalance(number int) {
	account.RWMutex.Lock()
	account.balance = account.balance + number
	account.RWMutex.Unlock()
}

func (account *Account) getBalance() int {
	account.RWMutex.RLock()
	value := account.balance
	account.RWMutex.RUnlock()

	return value
}

func TestRWmutex(t *testing.T) {
	account := Account{}
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				account.addBalance(1)
				fmt.Println(account.getBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("total balance ", account.getBalance())

}
