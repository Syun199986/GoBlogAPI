package handlers

import (
	"encoding/json"
	// "errors"
	"fmt"
	// "io"
	"net/http"
	"strconv"

	"github.com/Syun199986/GoBlogAPI/models"
	"github.com/gorilla/mux"
	// "github.com/sqs/goreturns/returns"
)

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	article := reqArticle
	json.NewEncoder(w).Encode(article)
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int
	// "page"というキーが存在し、かつpageのValueが1つ以上存在するとき
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		// pageのValueが複数でも1つ目の値を採用
		page, err = strconv.Atoi(p[0])
		// エラーなら400番エラーを返す
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
		// "page"というキーが存在しなければpageは1にする
	} else {
		page = 1
	}

	// エラー回避用
	fmt.Println(page)

	article := []models.Article{models.Article1, models.Article2}
	json.NewEncoder(w).Encode(article)
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	// エラー回避用
	fmt.Println(articleID)

	article := models.Article1
	json.NewEncoder(w).Encode(article)
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	article := models.Article1
	json.NewEncoder(w).Encode(article)
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	comment := models.Comment1
	json.NewEncoder(w).Encode(comment)
}
