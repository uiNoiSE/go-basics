package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
)

type account struct {
	login    string
	password string
	url      string
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

func newAccount(login, password, urlString string) (*account, error) {

	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		fmt.Println("Неверный формат url")
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &account{
		login:    login,
		password: password,
		url:      urlString,
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

	myAccount, err := newAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный логин или пароль")
		return
	}

	myAccount.outputPassword()
}
