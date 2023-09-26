package postgresql

import (
	"database/sql"
	"doslab/internal/config"
	"fmt"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitializeDb() {
	var err error

	connStr := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", config.Cfg.DataBaseUser, config.Cfg.DataBasePassword, config.Cfg.DataBaseName)
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func Migrations() error {
	sqlCommand := "CREATE TABLE IF NOT EXISTS statistics (id serial,post_id integer,word text,count integer);"
	_, err := DB.Query(sqlCommand)
	return err
}
