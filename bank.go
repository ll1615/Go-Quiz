// keep unsafe variable in multi-goroutine in a single coroutine
// to secure the variable.

package bank

import "fmt"

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balances

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	if Balance() < amount {
		return false
	}
	deposits <- -amount
	return true
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}
