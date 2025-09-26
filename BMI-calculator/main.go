package main

import (
	"errors"
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recover ", r)
		}
	}()
	fmt.Println("__ Калькулятор индекса массы тела __")
	for {
		userHeight, userWeight := getUserInput()
		IMT, err := calculateIMT(userHeight, userWeight)
		if err != nil {
			// fmt.Println("Не заданы параметры для расчёта")
			// continue
			panic("Не заданы параметры для расчёта")
		}
		outputResult(IMT)
		isRepeateCalculation := checkRepeatCalculation()
		if !isRepeateCalculation {
			break
		}
	}
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
	fmt.Println(result)
	switch {
	case imt < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case imt < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case imt < 25:
		fmt.Println("У вас нормальный вес")
	case imt < 30:
		fmt.Println("У вас избыточный вес")
	case imt < 35:
		fmt.Println("У вас 1-я степень ожирения")
	case imt < 40:
		fmt.Println("У вас 2-я степень ожирения")
	default:
		fmt.Println("У вас степень ожирения")
	}
}

func calculateIMT(userHeight float64, userWeight float64) (float64, error) {
	if userHeight <= 0 || userWeight <= 0 {
		return 0, errors.New("не заданы рост или вес")
	}
	IMT := userWeight / math.Pow(userHeight/100, IMTPower)
	return IMT, nil
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userWeight float64
	fmt.Print("Введите свой рост в сантиментрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userWeight)
	return userHeight, userWeight
}

func checkRepeatCalculation() bool {
	var userChoise string
	fmt.Print("Вы хотите сделать ещё расчёт (y/n): ")
	fmt.Scan(&userChoise)
	return userChoise == "y"
}
