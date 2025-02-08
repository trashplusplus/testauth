package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	connStr := "user=postgres dbname=testauth password=admin sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	sqlBytes, err := os.ReadFile("../init.sql")
	if err != nil {
		return nil, fmt.Errorf("не вдалося прочитати init.sql: %w", err)
	}

	sqlStatements := strings.Split(string(sqlBytes), ";")
	for _, stmt := range sqlStatements {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}
		_, err := db.Exec(stmt)
		if err != nil {
			return nil, fmt.Errorf("помилка виконання запиту: %s, %w", stmt, err)
		}
	}

	fmt.Println("init.sql успішно виконано")

	log.Println("Database connected")
	return db, err

}
