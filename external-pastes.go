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
	resp, _ := http.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%s", id))

	var result externalPaste
	bytes, _ := io.ReadAll(resp.Body)
	json.Unmarshal(bytes, &result)

	return result
}
