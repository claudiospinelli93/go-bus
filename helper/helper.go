package helper

import (
	"fmt"
	"os"
	"strconv"
)

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func PromptUser(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scanln(&input)
	return input
}

func PromptUserInt(prompt string, defaultValue int) (int, error) {
	var enterMaxMessagesCount string
	var maxMessagesCount int
	var err error
	interation := 0
	for {
		enterMaxMessagesCount = PromptUser(prompt)
		if enterMaxMessagesCount == "" {
			maxMessagesCount = defaultValue
			break
		}

		maxMessagesCount, err = strconv.Atoi(enterMaxMessagesCount)
		if err == nil {
			break
		}

		interation++
		fmt.Println("Invalid input. Please enter a valid integer.")

		if interation == 3 {
			fmt.Println("Max attempts reached.")
			break
		}
	}
	return maxMessagesCount, err
}
