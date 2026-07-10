package main

import (
	"errors"
	"fmt"
)

type User struct {
	ID      string
	Balance float64
	Name    string
}

func (user *User) deposit(sum float64) {
	user.Balance += sum
}

func (user *User) withdraw(sum float64) error {
	if user.Balance >= sum {
		user.Balance -= sum
		return nil
	}
	return errors.New("недостаточно средств на балансе")
}

func main() {
	a := &User{"1", 100, "vasya"}
	println(a.Balance)
	a.deposit(100)
	println(a.Balance)
	err := a.withdraw(50)
	if err != nil {
		fmt.Println(err)
	}
	println(a.Balance)
	err = a.withdraw(200)
	if err != nil {
		fmt.Println(err)
	}
	println(a.Balance)

	b := &User{"2", 1500, "nevasya"}
	println(b.Balance)
	b.deposit(600)
	println(b.Balance)
	err = b.withdraw(600)
	if err != nil {
		fmt.Println(err)
	}
	println(b.Balance)
}
