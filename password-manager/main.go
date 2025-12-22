package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

type account struct {
	login    string
	password string
	url      string
}

type accountWithTimestamp struct {
	account
	createdAt time.Time
	updatedAt time.Time
}

var letterRunes = []rune("abcdefghijklmnoprstuvwxyzABCDEFGHIJKLMNOPRSTUVWXYZ1234567890-*!")

func (acc *account) generatePassword(n int) {
	res := make([]rune, n)

	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}

	acc.password = string(res)
}

func promptData(prompt string) string {
	fmt.Print(prompt, ": ")

	var res string
	fmt.Scanln(&res)

	return res
}

// func newAccount(login, password, urlString string) (*account, error) {
//
// 	if login == "" {
// 		return nil, errors.New("INVALID_LOGIN")
// 	}
//
// 	_, err := url.ParseRequestURI(urlString)
// 	if err != nil {
// 		fmt.Println("Неверный формат url")
// 		return nil, errors.New("INVALID_URL")
// 	}
//
// 	newAcc := &account{
// 		login:    login,
// 		password: password,
// 		url:      urlString,
// 	}
//
// 	if password == "" {
// 		newAcc.generatePassword(12)
// 	}
//
// 	return newAcc, nil
// }

func newAccountWithTimestamp(login, password, urlString string) (*accountWithTimestamp, error) {

	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		fmt.Println("Неверный формат url")
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &accountWithTimestamp{
		account: account{
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

func (acc *account) outputPassword() {
	fmt.Println(acc.login, acc.password, acc.url)
}

func main() {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := newAccountWithTimestamp(login, password, url)
	if err != nil {
		fmt.Println("Неверный логин или пароль")
		return
	}

	fmt.Println(myAccount)
	myAccount.outputPassword()
}
