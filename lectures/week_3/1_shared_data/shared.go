// Go рутини, що одночасно працюють із загальними даними самі собою не можуть синхронізуватися
package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// Нехай у нас є Рахунок
type Account struct {
	balance float64
	mu      sync.RWMutex
}

func (a *Account) Balance() float64 {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.balance
}

func (a *Account) Deposit(amount float64) {
	log.Printf("depositing: %f (%f)", amount, a.Balance())
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance += amount
}

func (a *Account) Withdraw(amount float64) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if amount > a.balance {
		log.Printf("not enouth : %f", a.balance)
		return
	}
	log.Printf("withdrawing: %f (%f)", amount, a.balance)
	a.balance -= amount
}

func main() {
	acc := Account{}
	// Стартуємо 10 go-lessons рутин
	for i := 0; i < 10; i++ {
		go func(acc *Account, no int) {
			start := time.Now()
			// Кожна з яких проводить операції з акаунтом
			for j := 0; j < 10; j++ {
				// Іноді знімає гроші
				if j%2 == 0 {
					acc.Withdraw(50)
					continue
				}
				// іноді кладе
				acc.Deposit(50)
			}
			log.Printf("# %d completed in %v", no, time.Since(start))
		}(&acc, i)
	}

	_, _ = fmt.Scanln()
	// Що ж вийде в результаті
	fmt.Printf("balance: %f", acc.Balance())
}

// // //
// func closure() {
// // Функції захоплюють змінні у сфері видимості
// // Але щоб передавати значення, потрібно явно передавати їх у функцію
// for i: = 0; i < 10; i++ {
// go func() {
// fmt.Println("Got", i)
// } ()
// }
// fmt.Scanln()
// }
