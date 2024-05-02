package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
)

func main() {

	// Open or create the SQLite database file
	log.Printf("Opening database...")

	db, err := sql.Open("sqlite3", "posts.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Printf("Creating posts table...")
	// Create the posts table if it doesn't exist
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS posts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT,
			content TEXT
		);
	`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Hello World!")
	})

	// a route to create an anonymous post
	router.HandleFunc("POST /posts", func(w http.ResponseWriter, r *http.Request) {
		// Parse the request body
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			return
		}

		// Get the title and content from the form data
		title := r.Form.Get("title")
		content := r.Form.Get("content")

		if title == "" || content == "" {
			http.Error(w, "invalid: title and content are required", http.StatusBadRequest)
			return
		}

		log.Printf("title: %s content: %s", title, content)

		// Insert the post into the database
		_, err = db.Exec("INSERT INTO posts (title, content) VALUES (?, ?)", title, content)
		if err != nil {
			http.Error(w, "Failed to insert post into database", http.StatusInternalServerError)
			return
		}

		// Return a success response
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, "Post created")
	})

	// a route to get the most recent 10 posts
	router.HandleFunc("GET /posts/", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, title, content FROM posts ORDER BY id DESC LIMIT 10")
		if err != nil {
			http.Error(w, "Failed to retrieve posts from database", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		type Post struct {
			ID      int    `json:"id"`
			Title   string `json:"title"`
			Content string `json:"content"`
		}

		var posts []Post
		for rows.Next() {
			var post Post
			err := rows.Scan(&post.ID, &post.Title, &post.Content)
			if err != nil {
				http.Error(w, "Failed to retrieve posts from database", http.StatusInternalServerError)
				return
			}
			posts = append(posts, post)
		}

		jsonPosts, err := json.Marshal(posts)
		if err != nil {
			http.Error(w, "Failed to convert posts to JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonPosts)
	})

	log.Printf("Starting server on [::]:8080")
	http.ListenAndServe("[::]:8080", router)
}
