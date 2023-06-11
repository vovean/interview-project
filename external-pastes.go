package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type externalPaste struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func getExternalPastes(id string) externalPaste {
	resp, err := http.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%s", id))
	if err != nil {
		fmt.Println("cannot get external data")
		return externalPaste{}
	}

	var result externalPaste
	bytes, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(bytes, &result); err != nil {
		return externalPaste{}
	}

	return result
}
