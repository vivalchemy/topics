package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

type Bank []*BankAccount

type BankAccount struct {
	balance int
	lock    sync.RWMutex
}

func (b *BankAccount) Deposit(amount int) {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.balance += amount
}

func (b *BankAccount) Withdraw(amount int) {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.balance -= amount
}

func (b *BankAccount) Balance() int {
	b.lock.RLock()
	defer b.lock.RUnlock()
	return b.balance
}

func performBankOperation(bank *Bank, wg *sync.WaitGroup) {
	defer wg.Done()
	bankAccountIndex := rand.IntN(len(*bank))
	bankAccount := (*bank)[bankAccountIndex]

	for {
		time.Sleep(time.Duration(rand.IntN(1000)) * time.Millisecond)
		switch rand.IntN(3) {
		case 0:
			bankAccount.Deposit(rand.IntN(100))
			fmt.Println("Deposit in ", bankAccountIndex)
		case 1:
			bankAccount.Withdraw(rand.IntN(100))
			fmt.Println("Withdraw in ", bankAccountIndex)
		case 2:
			fmt.Println(bankAccount.Balance())
			fmt.Println("Balance in ", bankAccountIndex)
		}
	}
}

func main() {
	var bank Bank
	bank = append(bank, &BankAccount{balance: 100})
	bank = append(bank, &BankAccount{balance: 200})

	var wg sync.WaitGroup

	for range 10 {
		wg.Add(1)
		go performBankOperation(&bank, &wg)
	}

	wg.Wait()

}
