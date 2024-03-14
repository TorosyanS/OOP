package bank

import (
	"errors"
	"fmt"
)

type Bank struct {
	storage map[string]accountData
}

func NewBank() *Bank {
	return &Bank{
		storage: map[string]accountData{
			"100": {
				number:  "100",
				balance: 1000,
			},
		},
	}
}

func (b *Bank) TransferMoney(senderAccountNumber string, receiverAccountNumber string, amount int64) error {
	senderAccountData, err := b.FindAccount(senderAccountNumber)
	if err != nil {
		return fmt.Errorf("sender %v", err)
	}

	receiverAccountData, err := b.FindAccount(receiverAccountNumber)
	if err != nil {
		return fmt.Errorf("receiver %v", err)
	}

	if senderAccountData.balance-amount < 0 {
		return errors.New("not enough money for transfer")
	}

	b.ChangeBalance(senderAccountNumber, senderAccountData.balance-amount)
	b.ChangeBalance(receiverAccountNumber, receiverAccountData.balance+amount)

	return nil
}

func (b *Bank) FindAccount(accountNumber string) (accountData, error) {
	data, found := b.storage[accountNumber]
	if !found {
		return data, errors.New("account not found")
	}
	return data, nil
}

func (b *Bank) CreateNewAccount(number string) {
	b.storage[number] = accountData{
		number: number,
	}
}

func (b *Bank) ChangeBalance(accountNumber string, newBalance int64) {
	data := b.storage[accountNumber]
	data.balance = newBalance
	b.storage[accountNumber] = data
}

type accountData struct {
	number  string
	balance int64
}

func (a *accountData) Number() string {
	return a.number
}

func (a *accountData) Balance() int64 {
	return a.balance
}
