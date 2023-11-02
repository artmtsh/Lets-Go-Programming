package main

import (
	"GO_PROJECT3/calculations"
	"flag"
	"fmt"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func main() {
	var logEnable bool
	flag.BoolVar(&logEnable, "log", false, "Enable logging (default: false)")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("Необходимо передать число в аргументах командной строки")
		return
	}
	num, err := strconv.Atoi(os.Args[len(os.Args)-1])
	if err != nil {
		fmt.Println("Ошибка при преобразовании в число:", err)
		return
	}

	factorial, err := calculations.Calculate(int64(num), logEnable)
	if err == nil {
		log.Info("Result: ", factorial)
	} else {
		fmt.Println("Error:", err)
	}
}
