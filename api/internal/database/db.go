package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	_ "github.com/go-sql-driver/mysql" // mysql driver
)

var (
	DB  *sql.DB
	err error
)

// initializes the database connection
func InitDB() {
	// retrive the environment variables
	DbUser := os.Getenv("MYSQL_USER")
	DbPass := os.Getenv("MYSQL_ROOT_PASSWORD")
	DbHost := os.Getenv("MYSQL_HOST")
	DbPort := os.Getenv("MYSQL_PORT")
	DbName := os.Getenv("MYSQL_DATABASE")
	// database connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("MYSQL_USER"),          // db username env
		os.Getenv("MYSQL_ROOT_PASSWORD"), // db password env
		os.Getenv("MYSQL_HOST"),          // db host env
		os.Getenv("MYSQL_PORT"),          // db port env
		os.Getenv("MYSQL_DATABASE"),      // db name env
	)

	// print envs
	fmt.Printf("\n\033[1;1;96m * * * ⏳ Establishing connection to the database...\033[0m\n")
	fmt.Printf("\n\033[1;1;96m * * * 🛠️  Environment variables printed from \033[1;97;1mdb.go:\033[0m\n\n")
	fmt.Printf("\033[1;96m   User:          \033[1;97;1m➮ %s\033[0m\n", DbUser)
	fmt.Printf("\033[1;96m   Password:      \033[1;97;1m➮ %s*pass*%s \033[0m\n", string(DbPass[0]), string(DbPass[len(DbPass)-1]))
	fmt.Printf("\033[1;96m   Host:          \033[1;97;1m➮ %s\033[0m\n", DbHost)
	fmt.Printf("\033[1;96m   Port:          \033[1;97;1m➮ %s\033[0m\n", DbPort)
	fmt.Printf("\033[1;96m   Database Name: \033[1;97;1m➮ %s\033[0m\n", DbName)

	// mask the password
	MskPass := fmt.Sprintf("%s*pass*%s", string(DbPass[0]), string(DbPass[len(DbPass)-1]))
	// format the mkdsn
	MkDsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", DbUser, MskPass, DbHost, DbPort, DbName, dsn[strings.Index(dsn, "?")+1:])
	fmt.Printf("\n\033[1;96m   DSN:	      \033[1;96;1m➮ %s\033[0m\n\n", MkDsn)

	// format the dsn

	// connect to the database
	for {
		DB, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Printf("\033[1;31;1m 	* Failed to connect to the database: %v\033[0m", err)
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}

	// ping the database
	fmt.Println("\033[1;96;1m * * * 📡 Pinging DB...\033[0m")
	err = DB.Ping()
	if err != nil {
		fmt.Printf("\033[91m	* Error pinging database: %v\033[0m\n", err)
	}

	log.Println("\n\033[1;96;1m * * * ✅ Connected to database    at \033[1;97;1mhost: ➮", os.Getenv("MYSQL_HOST"), "\033[0m")
}

// closes the database connection
func CloseDB() {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			log.Printf("\033[1;91;1m	* Failed to close the database connection: %v\033[0m", err)
		} else {
			log.Println("\033[1;96;1m * * * Database connection closed successfully\033[0m")
		}
	}
}
