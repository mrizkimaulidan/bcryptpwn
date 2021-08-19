package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/fatih/color"
	"golang.org/x/crypto/bcrypt"
)

var (
	red    = color.New(color.FgRed).PrintlnFunc()
	green  = color.New(color.FgHiGreen).PrintlnFunc()
	yellow = color.New(color.FgYellow).PrintlnFunc()

	hashCommandArgument = os.Args[1]
	pathCommandArgument = os.Args[2]
)

func main() {
	file, err := os.Open(pathCommandArgument)

	if err != nil {
		panic("File not fucking found man!")
	}

	scanner := bufio.NewScanner(file)

	defer file.Close()

	for scanner.Scan() {
		// Remove all fucking whitespaces at the end of the sentence based on your wordlist
		sentence := strings.TrimSpace(scanner.Text())

		err := bcrypt.CompareHashAndPassword([]byte(hashCommandArgument), []byte(sentence))

		if err != nil {
			red("BAD ===>", sentence)
		}

		if err == nil {
			green("\nJACKPOT BITCH! ===>", sentence)
			green("\nI found it! Here is the plain hashes : " + "[ " + sentence + " ]")

			os.Exit(0)
		}
	}

	yellow("\nYour wordlist are trash! I cannot fucking find the hashes on your wordlist!")

	os.Exit(1)
}
