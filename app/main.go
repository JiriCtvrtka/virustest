package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

type virusChunk struct {
	Data string
}

func main() {
	endpoint := os.Getenv("ENDPOINT")
	if endpoint == "" {
		log.Panic("ENDPOINT env variable is not defined")
	}

	data := virusChunk{
		Data: "xxxxxx",
	}

	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	reader := bytes.NewReader(b)

	res, err := http.Post(endpoint, "application/json", reader)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func showdown() {
	cmd := exec.Command("shutdown", "/s", "/t", "0")
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}
