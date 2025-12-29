package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

type Account struct {
	login    string
	password string
	url      string
}

type AccountWithTimestamp struct {
	Account
	createdAt time.Time
	updatedAt time.Time
}

var letterRunes = []rune("abcdefghijklmnoprstuvwxyzABCDEFGHIJKLMNOPRSTUVWXYZ1234567890-*!")

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)

	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}

	acc.password = string(res)
}

func PromptData(prompt string) string {
	fmt.Print(prompt, ": ")

	var res string
	fmt.Scanln(&res)

	return res
}

func NewAccountWithTimestamp(login, password, urlString string) (*AccountWithTimestamp, error) {

	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		fmt.Println("Неверный формат url")
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &AccountWithTimestamp{
		Account: Account{
			login:    login,
			password: password,
			url:      urlString,
		},
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}

	if password == "" {
		newAcc.generatePassword(12)
	}

	return newAcc, nil
}

func (acc *Account) OutputPassword() {
	fmt.Println(acc.login, acc.password, acc.url)
}
