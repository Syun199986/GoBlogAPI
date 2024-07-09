package main

// import (
// 	"log"
// 	"net/http"

// 	"github.com/Syun199986/GoBlogAPI/handlers"
// 	"github.com/gorilla/mux"
// )

// func main() {
// 	r := mux.NewRouter()

// 	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
// 	r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
// 	r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
// 	r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodGet)
// 	r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
// 	r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

// 	log.Println("server start at port 8080")
// 	log.Fatal(http.ListenAndServe("localhost:8080", r))
// }

import (
	"fmt"
	"time"
)

type Comment struct {
	CommentID int
	ArticleID int
	Message   string
	CreatedAt time.Time
}

type Article struct {
	ID          int
	Title       string
	Contents    string
	UserName    string
	NiceNum     int
	CommentList []Comment
	CreatedAt   time.Time
}

func main() {
	comment1 := Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "test comment1",
		CreatedAt: time.Now(),
	}
	comment2 := Comment{
		CommentID: 2,
		ArticleID: 1,
		Message:   "second comment",
		CreatedAt: time.Now(),
	}
	article := Article{
		ID:          1,
		Title:       "first article",
		Contents:    "This is the test article",
		UserName:    "shun",
		NiceNum:     1,
		CommentList: []Comment{comment1, comment2},
		CreatedAt:   time.Now(),
	}
	fmt.Printf("%+v\n", article)
}
