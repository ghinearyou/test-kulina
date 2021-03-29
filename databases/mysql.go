package databases

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"

	"test-kulina/config"
	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	master *sqlx.DB
}

func New() *Mysql {
	config := config.Get()
	db, err := sqlx.Open("mysql",
		fmt.Sprintf("%s:%s@(%s:%s)/%s",
			config.DBUsername,
			config.DBPass,
			config.DBHost,
			config.DBPort,
			config.DBName,
		),
	)
	if err != nil {
		log.Fatalln(err)
	}

	return &Mysql {
		master: db,
	}
}

func (m *Mysql) GetActiveDB() *sqlx.DB {
	return m.master
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
