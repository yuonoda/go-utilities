package main

import (
	_ "github.com/lib/pq"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{})

	db.Create(&Product{Code: "D42", Price: 100})

	var product Product
	db.First(&product, 1)
	db.First(&product, "code = ?", "D42")

}

//type Person struct {
//	Name     string `json:"name"`
//	Nickname string `json:"nickname"`
//}
//
//const (
//	host     = "localhost"
//	port     = 5433
//	user     = "admin"
//	password = "admin"
//	dbname   = "postgres"
//)
//
//func OpenConnection() *sql.DB {
//	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
//		"password=%s dbname=%s sslmode=disable",
//		host, port, user, password, dbname)
//
//	db, err := sql.Open("postgres", psqlInfo)
//	if err != nil {
//		panic(err)
//	}
//
//	err = db.Ping()
//	if err != nil {
//		panic(err)
//	}
//
//	return db
//}
//
//func GETHandler(w http.ResponseWriter, r *http.Request) {
//	b, err := json.Marshal(r)
//	if err != nil {
//		http.Error(w, "処理に失敗しました", http.StatusInternalServerError)
//	}
//
//	log.Println("b:", b)
//	log.Println("GETHandler")
//	db := OpenConnection()
//
//	rows, err := db.Query("SELECT * FROM person")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	var people []Person
//
//	for rows.Next() {
//		var person Person
//		rows.Scan(&person.Name, &person.Nickname)
//		people = append(people, person)
//	}
//
//	peopleBytes, _ := json.MarshalIndent(people, "", "\t")
//
//	w.Header().Set("Content-Type", "application/json")
//	w.Write(peopleBytes)
//
//	defer rows.Close()
//	defer db.Close()
//}
//
//func POSTHandler(w http.ResponseWriter, r *http.Request) {
//	db := OpenConnection()
//
//	var p Person
//	err := json.NewDecoder(r.Body).Decode(&p)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	sqlStatement := `INSERT INTO person (name, nickname) VALUES ($1, $2)`
//	_, err = db.Exec(sqlStatement, p.Name, p.Nickname)
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		panic(err)
//	}
//
//	w.WriteHeader(http.StatusOK)
//	defer db.Close()
//}
//
//func main() {
//	http.HandleFunc("/", GETHandler)
//	http.HandleFunc("/insert", POSTHandler)
//	log.Fatal(http.ListenAndServe(":8080", nil))
//}
