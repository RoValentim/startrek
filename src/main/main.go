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

	var hex, specie, word string

	argsCount := len(os.Args)-1
	argsData  := os.Args[1:]

	if argsCount < 1 {
		logger.Log("0", "No word to translate", logger.ERROR)
		return
	}

	for i, v := range argsData {
		word = word + " " + v
		logger.Log(strconv.Itoa(i), v, logger.DEBUG)
	}
	word = strings.Trim(word, " ")

	translate.SetpIqaD()

	ch := make(chan bool, 2)
	go translate.Klingon  (ch, word, &hex   )
	go character.GetSpecie(ch, word, &specie)
	<- ch
	<- ch
	close(ch)

	logger.Log("0", hex   , logger.DEFAULT)
	logger.Log("0", specie, logger.DEFAULT)

	logger.Log("0", "Ending Star Trek Transaltor", logger.FATAL)
}
