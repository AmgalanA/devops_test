package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	_ "github.com/lib/pq"
)

const (
	host = "postgres"
	port = 5432
	user = "postgres"
	password = "postgres"
	dbname = "mydb"
)

func main() {
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)
	db, err := sql.Open("postgres", dbInfo)
	
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	http.HandleFunc("/save", func(w http.ResponseWriter, r *http.Request) {
		_, err := db.Exec("CREATE TABLE IF NOT EXISTS messages (content TEXT)")
		
		if err != nil {
			log.Fatal(err)
		}

		_, err = db.Exec("INSERT INTO messages (content) VALUES ('Hello from Docker')")

		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintln(w, "Message saved!")
	})

	http.HandleFunc("/read", func(w http.ResponseWriter, r *http.Request) {
		var content string
		
		row := db.QueryRow("SELECT content FROM messages LIMIT 1")

		err := row.Scan(&content)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintln(w, "Message from DB: %s", content)
	})

	log.Println("Server is starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
