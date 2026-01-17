package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	count := 20
	wrongAnswers := 0
	thinkingTime := 0

	fmt.Printf("%d Rechenaufgaben\n", count)
	fmt.Printf("Höchste Zahl? ")
	maxNumber := readInteger()

	fmt.Printf("Los geht's!\n")
	for range count {
		_wrongAnswers, _thinkingTime := giveMathProblem(maxNumber)
		wrongAnswers += _wrongAnswers
		thinkingTime += _thinkingTime
	}

	showStatistics(count, wrongAnswers, float64(thinkingTime))
}

func giveMathProblem(maxNumber int) (int, int) {
	wrongAnswers := 0
	thinkingTime := 0

	number1 := maxNumber/2 + rand.Intn(maxNumber/2+1)
	number2 := maxNumber/2 + rand.Intn(maxNumber/2+1)
	result := number1 + number2
	fmt.Printf("Wie viel ist %d + %d = ", number1, number2)
	for {
		start := time.Now().UnixMilli()
		answer := readInteger()
		end := time.Now().UnixMilli()
		thinkingTime += (int)(end - start)
		if answer == result {
			fmt.Printf("Richtig!\n")
			return wrongAnswers, thinkingTime
		}
		fmt.Printf("Falsch! %d + %d = ", number1, number2)
		wrongAnswers++
	}
}

func showStatistics(count, wrongAnswers int, thinkingTime float64) {
	if wrongAnswers > 0 {
		fmt.Printf("Du hast %d mal falsch geantwortet und ", wrongAnswers)
	} else {
		fmt.Printf("Alles richtig! Du ")
	}
	fmt.Printf("hast für %d Rechenaufgaben %.1f Sekunden nachgedacht!\n", count, float64(thinkingTime)/1000.0)
}

func readInteger() int {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		value, err := strconv.Atoi(input)
		if err == nil {
			return value
		}
	}
}
