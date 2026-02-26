package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/pflag"
)

func main() {
	count := pflag.IntP("count", "c", 20, "Number of problems to generate and solve.")
	limit := pflag.IntP("limit", "l", 100, "Maximum number used in generated problems.")
	pflag.Parse()

	mistakes := 0
	startTime := time.Now().UnixMilli()
	for range *count {
		handleTask(*limit, &mistakes)
	}
	endTime := time.Now().UnixMilli()

	fmt.Printf("You solved %d tasks in %d seconds!\n", *count, (endTime-startTime)/1000)
	fmt.Printf("You have made %d mistakes.\n", mistakes)
}

func handleTask(limit int, mistakes *int) {
	task := NewTask(limit)
	for {
		fmt.Print(task.Challenge())
		answer := readInteger()
		if answer == task.Result() {
			return
		}
		fmt.Println("Wrong!!! Repeat ...")
		*mistakes++
	}
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
