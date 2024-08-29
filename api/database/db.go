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
	os.Getenv("MYSQL_USER"),     // db username env
	os.Getenv("MYSQL_ROOT_PASSWORD"), // db password env
	os.Getenv("MYSQL_HOST"),     // db host env
	os.Getenv("MYSQL_PORT"),     // db port env
	os.Getenv("MYSQL_DATABASE"),     // db name env
)

	// print envs
	fmt.Printf("\n\033[1;34m * * * Establishing connection to the database...\033[0m")
	fmt.Printf("\n\033[1;37m * Environment variables print from \033[1;36mdb.go:\n\n\033[1;36m")
	fmt.Printf("   User:          ➮ %s\n", DbUser)
	fmt.Printf("   Password:      ➮ %s*pass*%s \n", string(DbPass[0]), string(DbPass[len(DbPass)-1]))
	fmt.Printf("   Host:          ➮ %s\n", DbHost)
	fmt.Printf("   Port:          ➮ %s\n", DbPort)
	fmt.Printf("   Database Name: ➮ %s\n", DbName)

	// mask the password
    MskPass := fmt.Sprintf("%s*pass*%s", string(DbPass[0]), string(DbPass[len(DbPass)-1]))
    // format the mkdsn
    MkDsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", DbUser, MskPass, DbHost, DbPort, DbName, dsn[strings.Index(dsn, "?")+1:])
	fmt.Printf("   DSN:		  ➮ %s\033[0m\n\n", MkDsn)


	// format the dsn
	

	// connect to the database
	for {
		DB, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Printf("\033[1;31;1m * Failed to connect to the database: %v\033[0m", err)
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}

	// ping the database
	fmt.Println(" * Pinging DB...")
	err = DB.Ping()
	if err != nil {
		fmt.Printf("\033[31m	Error pinging database: %v\033[0m\n", err)
	}

	log.Println("\033[1;37;1m * Connected to database at address: ➮\033[1;94;1m", os.Getenv("MYSQL_HOST"), "\033[0m")
}

// closes the database connection
func CloseDB() {
    if DB != nil {
        err := DB.Close()
        if err != nil {
            log.Printf("\033[1;31;1m * Failed to close the database connection: %v\033[0m", err)
        } else {
            log.Println("\033[1;37;1m * Database connection closed successfully\033[0m")
        }
    }
}