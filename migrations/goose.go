// This is custom goose binary with sqlite3 support only.

package migrations

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"
	"github.com/spf13/viper"
)

func Migration() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)

	var db *sql.DB
	// setup database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	db, _ = sql.Open("mysql", dsn)

	if err := db.Ping(); err != nil {
		log.Fatalln(string("\033[31m"), "error connection: ", err.Error())
		return
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	if err := goose.SetDialect("mysql"); err != nil {
		panic(err)
	}

	s, _ := goose.GetDBVersion(db)
	fmt.Println("version of db", s)

	if err := goose.Run("up", db, "migrations"); err != nil {
		panic(err)
	}
}
