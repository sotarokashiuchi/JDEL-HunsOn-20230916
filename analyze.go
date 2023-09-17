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
	if number%2 == 0 {
		fmt.Println(number, "は偶数です")
	}
	return
}

func isPrimeNumber(number int) {
	var i int
	if number <= 1 {
		return
	}
	for i = 2; i < number; i++ {
		if number%i == 0 {
			return
		}
	}
	fmt.Println("素数です！")
	return
}

func isPerfectNumber(number int) {
	var i, sum int
	for i = 1; i < number; i++ {
		if number%i == 0 {
			sum += i
		}
	}
	if sum == number {
		fmt.Println(number, "は完全数でえす")
	}
	return
}

func isAmicableNumbers(x, y int) {
	var i int
	var sum int
	var sum3 int
	for i = 1; i < x; i++ {
		if x%i == 0 {
			sum = sum + i
		}
	}
	if sum == y {
		for i = 1; i < y; i++ {
			if y%i == 0 {
				sum3 = sum3 + i
			}
		}
		if sum3 == x {
			fmt.Println("親和数です！")
		}

	}
	return
}
