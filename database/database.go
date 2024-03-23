package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct{}

var DB *sql.DB

func Connect() {
	connection := "dev:dev@tcp(mysql8:3306)/management-personal"
	db, err := sql.Open("mysql", connection)

	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()

	if err != nil {
		panic(fmt.Sprintf("Erro ao conectar com o banco de dados: %s", err))
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	fmt.Println("Conexão com o banco de dados estabelecida")
	DB = db
}
