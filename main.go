package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	target := rand.Intn(100) + 1

	// fmt.Println(target)
	fmt.Println("Я загадал случайное число от 1 до 100")
	fmt.Println("Угадаешь его?")

	reader := bufio.NewReader(os.Stdin)
	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("У тебя осталось ", 10-guesses, " попыток")

		fmt.Println("Введи новое число: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)

		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Ой. Загаданное число больше введенного")
		} else if guess > target {
			fmt.Println("Ой. Загаданное число меньше введенного")
		} else {
			fmt.Println("Ура! Загаданое число угадано!!!")
			success = true
			break
		}
		// fmt.Println(target)
	}
	if !success {
		fmt.Println("Жалко, ты не смог угадать число... :-(")
	}
}
