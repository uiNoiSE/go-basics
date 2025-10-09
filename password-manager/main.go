package main

import (
	"fmt"
	"math/rand/v2"
)

type account struct {
	login    string
	password string
	url      string
}

func (acc account) outputPassword() {
	fmt.Println(acc.login, acc.password, acc.url)
}

var letterRunes = []rune("abcdefghijklmnoprstuvwxyzABCDEFGHIJKLMNOPRSTUVWXYZ1234567890-*!")

func (acc *account) generatePassword(n int) {
	res := make([]rune, n)

	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}

	acc.password = string(res)
}

func main() {
	login := promptData("Введите логин")
	url := promptData("Введите URL")

	myAccount := account{
		login: login,
		url:   url,
	}

	myAccount.generatePassword(12)
	myAccount.outputPassword()
}

func promptData(prompt string) string {
	fmt.Print(prompt, ": ")

	var res string
	fmt.Scan(&res)

	return res
}
