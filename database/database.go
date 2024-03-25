package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"management-personal/config"
	"management-personal/utils"
)

type Database struct {
	DB *sql.DB
}

func Connect(d *Database) {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%d)/management-personal", config.Env.DbUser, config.Env.DbPassword, config.Env.DbHost, config.Env.DbPort)
	db, err := sql.Open("mysql", connection)

	utils.NewError(err)

	err = db.Ping()

	utils.NewError(err)

	fmt.Println("Conexão com o banco de dados estabelecida")
	db = d.DB
}
