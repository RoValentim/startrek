package messages

import (
	"testing"
)

func TestNotFound_pass(t *testing.T) {
        if len(ReturnList) > 3 {
                t.Errorf("Finding unset message")
        }
}

func TestFound_pass(t *testing.T) {
        if ReturnList[0].Message == "" {
                t.Errorf("Error finding message")
        }
}
