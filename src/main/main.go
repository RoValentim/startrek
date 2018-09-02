package main

import (
	"os"
	"strconv"
	"strings"

	"services/character"
	"services/translate"

        "shared/logger"
)

func main() {
	logger.Log("0", "Starting Star Trek Transaltor", logger.FATAL)

	var word string

	argsCount := len(os.Args)-1
	argsData  := os.Args[1:]

	if argsCount < 1 {
		logger.Log("0", "No word to translate", logger.ERROR)
		return
	}

	ch := make(chan bool, 1)
	translate.SetpIqaD(ch)

	for i, v := range argsData {
		word = word + " " + v
		logger.Log(strconv.Itoa(i), v, logger.DEBUG)
	}
	word = strings.Trim(word, " ")

	<- ch
        close(ch)

	hex := translate.Klingon(word)
	logger.Log("0", hex, logger.DEFAULT)

	specie := character.GetSpecie(word)
	logger.Log("0", specie, logger.DEFAULT)

	logger.Log("0", "Ending Star Trek Transaltor", logger.FATAL)
}
