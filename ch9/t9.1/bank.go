package bank

import "sync"

type depositsSync struct {
	deposits chan int // send amount to deposit
	sync.RWMutex
}

var deposits = depositsSync{deposits: make(chan int)}
var balances = make(chan int) // receive balance

func Deposit(amount int) {
	deposits.Lock()
	defer deposits.Unlock()
	deposits.deposits <- amount
}
func Balance() int { return <-balances }
func Withdraw(amount int) bool {
	deposits.RLock()
	if <-balances < amount {
		deposits.RUnlock()
		return false
	} else {
		deposits.RUnlock()
		deposits.Lock()
		defer deposits.Unlock()
		deposits.deposits <- -amount
		return true
	}
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits.deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-
