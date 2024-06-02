package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"golang-lsp/lsp"
	"golang-lsp/rpc"
	"log"
	"os"
)

func main() {
	logger := getLogger("/home/drama321/coding/golang-lsp/log.txt")
	logger.Println("Hey, I started!")

	scanner := bufio.NewScanner(os.Stdin) // read from stdin
	scanner.Split(rpc.Split)
	for scanner.Scan() { // scan till you can't no more
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Got an error: %s", err)
		}
		handleMessage(logger, method, contents)

	}
	fmt.Println("hello world!")
}

func handleMessage(logger *log.Logger, method string, contents []byte) {
	logger.Printf("Recievied msg with method: %s", method)
	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Could not parse this: %s", err)
		}
		logger.Printf("Connected to: %s %s", request.Params.ClientInfo.Name, request.Params.ClientInfo.Version)
		msg := lsp.NewInitializeResponse(request.ID)
		reply := rpc.EncodeMessage(msg)
		writer := os.Stdout
		writer.Write([]byte(reply))
		logger.Println("Sent the reply Response")
	}
}

func getLogger(fileName string) *log.Logger {
	logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("hey, you didnt give me a valid file")
	}

	return log.New(logFile, "[educationlsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
