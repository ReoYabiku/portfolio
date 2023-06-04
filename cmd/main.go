package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"portfolio/app/handler"
	"portfolio/app/infra"
	"portfolio/app/middleware"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")

	time.Sleep(2 * time.Second)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, db_port, database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	mysql := infra.NewMysql(db)

	bh := handler.NewBiography(mysql)
	ah := handler.NewAchievement(mysql)

	http.HandleFunc("/biographies", middleware.CORS(bh.GetAll))
	http.HandleFunc("/achievements", middleware.CORS(ah.GetAll))

	fmt.Println("server is running...")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
