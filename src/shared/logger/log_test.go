package logger

import (
	"os"
	"testing"
)

func TestLog_pass(t *testing.T) {
	file := "/var/log/stt.log";

	if _, err := os.Stat(file); os.IsNotExist(err) {
		if _, err := os.Stat("stt.log"); err == nil {
			file = "stt.log"
		} else {
			file = ""
		}
	}

	if file != "" {
		err := os.Remove(file)
		if err != nil {
			t.Errorf("Delete old log file error")
		}
	}

	Log("0", "--- Log Test", DEFAULT)

	if _, err := os.Stat(file); os.IsNotExist(err) {
		if _, err := os.Stat("stt.log"); err == nil {
			file = "stt.log"
		} else {
			file = ""
		}
	}

	if _, err := os.Stat(file); os.IsNotExist(err) {
		t.Errorf("Create new log file error")
	}

	err := os.Remove(file)
	if err != nil {
		t.Errorf("Error deleting log file")
	}
}
