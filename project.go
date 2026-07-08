package main

import (
	"errors"
	"fmt"
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

func (PS *PaymentSystem) ProcessingTransactions(tr Transaction) {
	i, ok := PS.Users[tr.FromID]
	if !ok {
		return
	}
	j, ok := PS.Users[tr.ToID]
	if !ok {
		return
	}
	err := i.withdraw(tr.Amount)
	if err != nil {
		fmt.Println(err)
	} else {
		j.deposit(tr.Amount)
	}
	fmt.Println(j.Balance, i.Balance)
}
func (user *User) deposit(sum float64) {
	user.Balance += sum
}

func (user *User) withdraw(sum float64) error {
	if user.Balance >= sum {
		user.Balance -= sum
		return nil
	}
	return errors.New("insufficient funds")
}

func main() {
	a := User{"1", 100, "vasya"}
	b := User{"2", 200, "petya"}
	PS := PaymentSystem{make(map[string]*User), []Transaction{}}
	PS.AddUser(&a, &b)
	tr1 := Transaction{a.ID, b.ID, 17}
	tr2 := Transaction{a.ID, b.ID, 14}
	tr3 := Transaction{a.ID, b.ID, 30}
	tr4 := Transaction{b.ID, a.ID, 54}
	tr5 := Transaction{b.ID, a.ID, 12}
	PS.AddTransaction(tr1, tr2, tr3, tr4, tr5)
	for _, val := range PS.Transactions {
		PS.ProcessingTransactions(val)
		fmt.Println(a.Balance, b.Balance)
	}
	PS.Transactions = *new([]Transaction)
	for _, val := range PS.Transactions {
		PS.ProcessingTransactions(val)
		fmt.Println(a.Balance, b.Balance)
	}
}
