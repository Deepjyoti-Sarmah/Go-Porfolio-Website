package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/spf13/viper"
	_ "github.com/go-sql-driver/mysql"
)

func LoadEnvVariable(key string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatal(err)
	}

	return value
}

func Connection(name, email, subject, body string) {
	db_user := LoadEnvVariable("DB_USER")
	db_password := LoadEnvVariable("DB_PASSWORD")
	db_address := LoadEnvVariable("DB_ADDRESS")
	db_db := LoadEnvVariable("DB_DB")
	s := fmt.Sprintf("%s:%s@tcp{%s:3306}/%s", db_user, db_password, db_address, db_db)
	db, err := sql.Open("mysql", s)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	fmt.Println("Database is successfully connected")
	CreateTable(db)
	InsertData(db, name, email, subject, body)
}

func CreateTable(db *sql.DB) {
	query, err := ioutil.ReadFile("database/person.SQL")
	if err != nil {
		log.Fatal(err)
	}

	request := strings.Split(string(query), ":")[0]

	stmt, err := db.Prepare(request)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		log.Fatal()
	}
}

func InsertData(db *sql.DB, name, email, subject, body string) {
	q := "INSERT INTO Person(name, email, subject, body) VALUES(?, ?, ?, ?)"
	insert, err := db.Prepare(q)
	if err != nil {
		log.Fatal(err)
	}
	defer insert.Close()

	_, err = insert.Exec(name, email, subject, body)
	if err != nil {
		log.Fatal(err)
	}
}
