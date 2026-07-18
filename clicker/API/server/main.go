package main

import (
	sql "database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"net/http"
	"fmt"
	"log"
	"io"
)

type s_dataBase struct {
	db *sql.DB
}

type User struct {
	ID			int
	Name		int
	money		int
};

func (dataBase *s_dataBase) getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getted")
	var (id int 
		login string
		money int)
	rows, err := dataBase.db.Query("SELECT * FROM player")
	if (err != nil) {
		fmt.Println("Query failed")
	} else {
		for rows.Next() {
		   rows.Scan(&id, &login, &money)
		   fmt.Printf("%d - %s - %s \n", id, login, money)
	  }
	}
	fmt.Println("bomb has been getted")
}

func (dataBase *s_dataBase)createUsers(w http.ResponseWriter, r *http.Request) {
	res, err := io.ReadAll(r.Body)
	fmt.Printf("request : %s\n",res)
	id := 0
	login := res
	money := 346
	_, err = dataBase.db.Exec("INSERT INTO player(id, login, money) VALUES($1, $2, $3)", id, login, money)
	if (err != nil) {
		fmt.Println(err)
		fmt.Println("db insert failed")
	} else {
		fmt.Println("bomb has been posted")
	}
}

func connectDataBase() (*s_dataBase, error){
	db, err := sql.Open("pgx", "postgres://user:pass@localhost:5432/clicker?sslmode=disable") 
	database := &s_dataBase{db : db}
	return database, err
}

func main () {
	dataBase, err := connectDataBase()
	if err != nil {
		fmt.Println("failed to create database")
	}
	router := http.NewServeMux()
	router.HandleFunc("GET /users", dataBase.getUsers)
	router.HandleFunc("POST /users", dataBase.createUsers)
	log.Fatal(http.ListenAndServe(":8080", router))
	dataBase.db.Close()
}
