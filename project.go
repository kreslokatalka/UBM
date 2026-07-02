package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)

type Transaction struct {
	FromID, ToID int
	Amount       float32
}
type User struct {
	ID      int
	Balance float32
	Name    string
	mu      sync.Mutex
}
type PaymentSystem struct {
	Users        map[int]*User
	Transactions []Transaction
}

func (PS *PaymentSystem) AddUser(us ...*User) {
	for _, u := range us {
		PS.Users[u.ID] = u
	}
}
func (PS *PaymentSystem) AddTransaction(tr ...Transaction) {
	PS.Transactions = append(PS.Transactions, tr...)
}

func (PS *PaymentSystem) ProcessingTransactions(tr Transaction) error {
	i, ok := PS.Users[tr.FromID]
	if !ok {
		return errors.New("Пользователь с таким id не найден")
	}
	j, ok := PS.Users[tr.ToID]
	if !ok {
		return errors.New("Пользователь с таким id не найден")
	}
	i.mu.Lock()
	j.mu.Lock()
	defer i.mu.Unlock()
	defer j.mu.Unlock()
	err := i.withdraw(tr.Amount)
	if err != nil {
		return err
	} else {
		j.deposit(tr.Amount)
		fmt.Println(i.Name, i.Balance, j.Name, j.Balance, tr.Amount)
		return nil
	}
}
func (user *User) deposit(sum float32) {
	user.Balance += sum
}

func (user *User) withdraw(sum float32) error {
	if user.Balance > sum {
		user.Balance -= sum
		return nil
	}
	return errors.New("insufficient funds")
}
func (ps *PaymentSystem) create_users(n int) {
	for i := 0; i < n; i++ {
		name := "user" + strconv.Itoa(len(ps.Users))
		us := &User{len(ps.Users), float32(rand.Intn(1000)), name, sync.Mutex{}}
		ps.AddUser(us)
	}
}
func (ps *PaymentSystem) create_tran(n int) error {
	if len(ps.Users) < 2 {
		return errors.New("недостаточно пользователей")
	}
	for i := 0; i < n; {
		from := rand.Intn(len(ps.Users))
		to := rand.Intn(len(ps.Users))
		if from == to {
			continue
		}
		amount := float32(rand.Intn(10000)) / 100
		tran := Transaction{from, to, amount}
		ps.AddTransaction(tran)
		i++
	}
	return nil
}
func (ps *PaymentSystem) worker(ch chan Transaction) {
	for tr := range ch {
		err := ps.ProcessingTransactions(tr)
		if err != nil {
			fmt.Printf("Ошибка транзакции: %v\n", err)
		}
	}
}
func main() {

	PS := PaymentSystem{make(map[int]*User), []Transaction{}}
	PS.create_users(10)
	PS.create_tran(100)
	Transaction_chan := make(chan Transaction, 100)
	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Go(func() { PS.worker(Transaction_chan) })
	}
	for i := range PS.Transactions {
		Transaction_chan <- PS.Transactions[i]
	}
	close(Transaction_chan)
	wg.Wait()
}
