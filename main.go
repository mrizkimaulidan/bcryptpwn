package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
	"golang.org/x/crypto/bcrypt"
)

var (
	red    = color.New(color.FgRed).PrintlnFunc()
	green  = color.New(color.FgHiGreen).PrintlnFunc()
	yellow = color.New(color.FgYellow).PrintlnFunc()
	blue   = color.New(color.FgBlue).PrintlnFunc()

	hashCommandArgument string
	pathCommandArgument string
)

func main() {
	flag.StringVar(&hashCommandArgument, "hash", "", "your bcrypt hashes")
	flag.StringVar(&pathCommandArgument, "path", "", "your wordlist path")
	flag.Parse()

	blue("Hash:", hashCommandArgument)

	file, err := os.Open(pathCommandArgument)
	if err != nil {
		log.Fatalln("failed opening path", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// Remove all fucking whitespaces at the end of the sentence based on your wordlist
		sentence := strings.TrimSpace(scanner.Text())
		err := bcrypt.CompareHashAndPassword([]byte(hashCommandArgument), []byte(sentence))
		if err != nil {
			red("BAD ===>", sentence)
		} else {
			green("\nJACKPOT BITCH! ===>", sentence)
			green("\nI found it! Here is the plain hashes : " + "[ " + sentence + " ]")

			os.Exit(0)
		}
	}

	yellow("Your wordlist are trash! I cannot fucking find right the hashes on your wordlist!")
}
