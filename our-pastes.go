package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/url"
	"strconv"
	"time"
)

type Paste struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt int64
	CreatorIP string
}

func getDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func (p Paste) saveToDB() int {
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	db := getDB()
	conn, _ := db.Conn(ctx)

	query := fmt.Sprintf("insert into pastes (title, body, created_at) values ('%s', '%s', %d)", p.Title, p.Body, time.Now().Unix())
	queryGetInsertedPaste := fmt.Sprintf("select id from pastes where title='%s'", p.Title)
	queryChngLog := "insert into changelog (paste_id, creator_ip, paste_body_len) values (%d, '%s', %d)"

	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	tx.Exec(query)
	insertedIdRes := tx.QueryRow(fmt.Sprintf(queryGetInsertedPaste))
	var insertedId int
	insertedIdRes.Scan(&insertedId)
	tx.Exec(fmt.Sprintf(queryChngLog, insertedId, p.CreatorIP, len(p.Body)))

	if err := tx.Commit(); err != nil {
		return -1
	}

	return insertedId
}

func getPasteFromDb(params url.Values) (Paste, bool) {
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	db := getDB()
	conn, _ := db.Conn(ctx)

	var query string

	id, err := strconv.Atoi(params.Get("id"))
	if err != nil {
		query = fmt.Sprintf("select title, body from pastes where title like '%s'", params.Get("title"))
	} else {
		query = fmt.Sprintf("select title, body from pastes where id=%d", id)
	}

	res := conn.QueryRowContext(ctx, query)
	var p Paste
	if err := res.Scan(&p.Title, &p.Body); err != nil {
		return p, false
	}
	return p, true
}
