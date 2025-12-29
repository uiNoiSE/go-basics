package main

import (
	"fmt"
	"go-basics/password-manager/account"
)

func main() {
	login := account.PromptData("Введите логин")
	password := account.PromptData("Введите пароль")
	url := account.PromptData("Введите URL")

	myAccount, err := account.NewAccountWithTimestamp(login, password, url)
	if err != nil {
		fmt.Println("Неверный логин или пароль")
		return
	}

	fmt.Println(myAccount)
	myAccount.OutputPassword()
}
