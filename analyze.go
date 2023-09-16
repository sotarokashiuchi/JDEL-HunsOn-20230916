package main

import "fmt"

func main() {
	var inputNumber []int
	var inputNumberTotal int

	// input
	fmt.Printf("解析する整数の個数を入力してください>>>")
	fmt.Scan(&inputNumberTotal)
	for i := 0; i < inputNumberTotal; i++ {
		var input int
		fmt.Printf("%vつ目の整数を入力してください>>>", i+1)
		fmt.Scan(&input)
		inputNumber = append(inputNumber, input)
	}

	fmt.Println("")
	fmt.Println("Analyze....")
	switch inputNumberTotal {
	case 1:
		isOddNumber(inputNumber[0])
		isEvenNumber(inputNumber[0])
		isPrimeNumber(inputNumber[0])
		isPerfectNumber(inputNumber[0])
	case 2:
		isAmicableNumbers(inputNumber[0], inputNumber[1])

	}
}

func isOddNumber(number int) {
	if number%2 == 1 {
		fmt.Println(number, "は奇数です")
	}
	return
}

func isEvenNumber(number int) {
	return
}

func isPrimeNumber(number int) {
	var i int
	for i = 2; i < number; i++ {
		if number%i == 0 {
			return
		}
	}
	fmt.Println("素数です！")
	return
}

func isPerfectNumber(number int) {
	return
}

func isAmicableNumbers(x, y int) {
	return
}
