package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"sync"
)

type Transaction struct {
	FromID, ToID string
	Amount       float64
}
type User struct {
	ID      string
	Balance float64
	Name    string
}
type PaymentSystem struct {
	Users        map[string]*User
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
	err := i.withdraw(tr.Amount)
	if err != nil {
		return err
	} else {
		j.deposit(tr.Amount)
		fmt.Println(i.Name, math.Round(i.Balance*100)/100, j.Name, math.Round(j.Balance*100)/100, tr.Amount)
		return nil
	}
}
func (user *User) deposit(sum float64) {
	user.Balance += sum
}

func (user *User) withdraw(sum float64) error {
	if user.Balance > sum {
		user.Balance -= sum
		return nil
	}
	return errors.New("insufficient funds")
}
func (ps *PaymentSystem) create_users(n int) {
	for i := 0; i < n; i++ {
		name := "user" + strconv.Itoa(len(ps.Users))
		us := &User{string(len(ps.Users)), float64(rand.Intn(1000)), name}
		ps.AddUser(us)
	}
}
func (ps *PaymentSystem) rand_user() string {
	keys := make([]string, 0, len(ps.Users))
	for key := range ps.Users {
		keys = append(keys, key)
	}
	return keys[rand.Intn(len(keys))]
}
func (ps *PaymentSystem) create_tran(n int) error {
	if len(ps.Users) < 2 {
		return errors.New("недостаточно пользователей")
	}

	for i := 0; i < n; {
		from := ps.rand_user()
		to := ps.rand_user()
		if from == to {
			continue
		}
		amount := float64(rand.Intn(10000)) / 100
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
			fmt.Println(tr.FromID, tr.ToID)
		}
	}

}
func main() {

	PS := PaymentSystem{make(map[string]*User), []Transaction{}}
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
