package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/external/", func(w http.ResponseWriter, r *http.Request) {
		externalId := strings.TrimPrefix("/external/", r.URL.Path) // тут баг - параметры нужно местами махнуть
		paste := getExternalPastes(externalId)

		marshelled, _ := json.Marshal(&paste)
		w.Write(marshelled)
		fmt.Println("returned external transaction success")
	})
	mux.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		var p Paste
		if err := json.Unmarshal(body, &p); err != nil {
			w.Write([]byte("400: bad request"))
			return
		}
		p.CreatedAt = time.Now().Unix()
		p.CreatorIP = r.RemoteAddr

		savedId := p.saveToDB()
		w.Write([]byte(fmt.Sprint(savedId)))
		fmt.Printf("saved new paste: %d", savedId)
	})
	mux.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		p, ok := getPasteFromDb(r.URL.Query())
		if !ok {
			fmt.Println("cannot get paste")
		}

		data, _ := json.Marshal(&p)
		w.Write(data)
	})

	http.ListenAndServe(":8080", mux)
}
