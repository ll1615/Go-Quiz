// keep unsafe variable in single coroutine
// to secure the variable.

package bank

type WithdrawData struct {
	amount   int
	resultCh chan bool
}

var withdrawCh = make(chan WithdrawData)
var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balances

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	ch := make(chan bool)
	withdrawCh <- WithdrawData{amount, ch}
	return <-ch
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case withdrawData := <-withdrawCh:
			if balance >= withdrawData.amount {
				balance -= withdrawData.amount
				withdrawData.resultCh <- true
			} else {
				withdrawData.resultCh <- false
			}
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}
