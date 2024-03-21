package main

import (
	"bufio"
	"log"
	"os"

	"github.com/brewinski/lsp-go/rpc"
)

func main() {
	logger := getLogger("/Users/chris.brewin/Documents/github/lsp-go/lsp.log")
	logger.Println("Starting LSP")

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Split(rpc.Split)

	for scanner.Scan() {
		handleMessage(logger, scanner.Text())
	}
}

func handleMessage(logger *log.Logger, msg any) {
	logger.Println("Received message: ", msg)
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("error opening log file: " + err.Error())
	}

	return log.New(logfile, "[brewinski:lsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
