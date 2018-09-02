package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"encoding/json"
)

const NONE    = 0
const DEFAULT = 1
const ERROR   = 2
const DEBUG   = 3
const FATAL   = 4

func Log(id string, data interface{}, level int) {
	ll, err := strconv.Atoi(os.Getenv("LOG_LEVEL"))

	if err != nil {
		ll = 1
	}

	if ll < level  {
		return
	}

	var logger  *log.Logger
	file := "/var/log/stt.log"

	fd, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		file := "stt.log"

		fd, err = os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Printf("Failed to open log file: %s : %#v\n", file, err)
			return
		}
	}
	defer fd.Close()

	destination := io.MultiWriter(fd, os.Stdout)

	switch e := level; e {
		case DEFAULT:
			logger = log.New(destination, "", 0)
		case ERROR:
			logger = log.New(destination, "[" + id + "] STT ERROR:   ", log.Ldate|log.Ltime)
		case DEBUG:
			logger = log.New(destination, "[" + id + "] STT DEBUG:   ", log.Ldate|log.Ltime)
		case FATAL:
			logger = log.New(destination, "[" + id + "] STT FATAL:   ", log.Ldate|log.Ltime)
	}

	var d map[string]interface{}

	m, err := json.Marshal(data)
	if err != nil {
		error := fmt.Sprintf("%v", data)
		logger.Print(error)
		return
        }

	err = json.Unmarshal(m, &d)
	if err != nil {
		error := fmt.Sprintf("%v", data)
		logger.Print(error)
		return
	}

	logger.Print(d)
}
