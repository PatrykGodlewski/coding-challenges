package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	buff, err := os.ReadFile("test.txt")
	if err != nil {
		t.Fail()
	}
	data := string(buff)
	if ReadWords(data) != "58164" {
		t.Fail()
	}

	if ReadLines(data) != "7145" {
		t.Fail()
	}

	if ReadChars(data) != "342190" {
		t.Fail()
	}

	if ReadRunes(data) != "339292" {
		t.Fail()
	}
}
