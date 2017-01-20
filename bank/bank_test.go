package bank

import (
	"fmt"
	"testing"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	for i := 0; i < 100; i++ {
		go func() {
			Deposit(200)

			if ok := Withdraw(20); ok {
				fmt.Println("Withdraw 20")
			} else {
				fmt.Println("Can't withdraw 20")
			}

			if ok := Withdraw(4000); ok {
				fmt.Println("Withdraw 4000")
			} else {
				fmt.Println("Can't withdraw 4000")
			}

			done <- struct{}{}
		}()
	}

	// Bob
	for i := 0; i < 100; i++ {
		go func() {
			Deposit(100)
			if ok := Withdraw(20); ok {
				fmt.Println("Withdraw 20")
			} else {
				fmt.Println("Can't withdraw 20")
			}

			if ok := Withdraw(4000); ok {
				fmt.Println("Withdraw 4000")
			} else {
				fmt.Println("Can't withdraw 4000")
			}
			done <- struct{}{}
		}()
	}

	// Wait for both transactions.
	for i := 0; i < 200; i++ {
		<-done
	}

	if got, want := Balance(), 0; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
