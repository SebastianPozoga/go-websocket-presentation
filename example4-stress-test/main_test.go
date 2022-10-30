package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	testDebbug := false
	testReaders := 1000
	testTimeout := 60 * 5
	testUrl := "wss://localhost:8080/ws"

	debbug = &testDebbug
	readerCounter = &testReaders
	timeout = &testTimeout
	url = &testUrl
	mainGo()
}
