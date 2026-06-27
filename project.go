package main

import (
	"errors"
	"fmt"
)

type User struct {
	ID      int
	Balance float32
	Name    string
}

func (user *User) deposit(sum float32) {
	user.Balance += sum
}

func (user *User) withdraw(sum float32) error {
	if user.Balance > sum {
		user.Balance -= sum
		return nil
	}
	return errors.New("недостаточно средств на балансе")
}

func main() {
	a := &User{1, 100, "vasya"}
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
}
