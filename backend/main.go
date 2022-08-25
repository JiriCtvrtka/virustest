package main

import (
	"fmt"
	"net/http"

	"github.com/percona/virustest/backend/storage"
)

func main() {
	http.HandleFunc("/save", storage.SaveData)

	if err := http.ListenAndServe(":8888", nil); err != nil {
		fmt.Print(err)
	}
}
