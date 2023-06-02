package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type biography struct {
	ID          int64
	Src         string
	Alt         string
	Summary     string
	Title       string
	Description string
}

func main() {
	bios, err := fetchSql()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "%v", bios)
	})

	fmt.Println("server is running...")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func fetchSql() ([]biography, error) {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")

	time.Sleep(2 * time.Second)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM biographies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bios := []biography{}

	for rows.Next() {
		b := &biography{}
		if err = rows.Scan(&b.ID, &b.Src, &b.Alt, &b.Summary, &b.Title, &b.Description); err != nil {
			return nil, err
		}
		bios = append(bios, *b)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return bios, err
}
