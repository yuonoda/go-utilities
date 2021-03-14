package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type User struct {
	UserId             int `gorm:"primaryKey"`
	UserFirstNameKanji string
	UserLastNameKanji  string
}

const (
	host     = "localhost"
	port     = 5433
	user     = "admin"
	password = "admin"
	dbname   = "postgres"
)

// TODO DBマイグレーションをバッチにする

func OpenDBConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Tokyo",
		host, user, password, dbname, port)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	db, err := OpenDBConnection()
	db.AutoMigrate(&User{})
	if err != nil {
		log.Println(fmt.Errorf("DB接続に失敗しました。%s", err))
	}
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var users []User
		db.Find(&users)
		log.Printf("%+v", users)
		usersBytes, _ := json.MarshalIndent(users, "", "\t")
		w.Write(usersBytes)

	} else if r.Method == "POST" {
		var u User
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		db.Create(&u)

	} else {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/users", usersHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
