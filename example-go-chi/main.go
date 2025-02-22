package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Article 结构体用于表示文章数据
type Article struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	// 1. Hello World
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	// 2. Demo for URL placeholder usage
	r.Get("/article/{id}", getArticle)

	// 3. Demo for get article list
	r.Post("/articles", getArticleList)

	http.ListenAndServe(":3000", r)
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	// 创建示例文章数据
	article := Article{
		ID:      id,
		Title:   "article title",
		Content: "This is post " + id + " content",
	}

	// 设置响应头为 JSON
	w.Header().Set("Content-Type", "application/json")

	// 将结构体转换为 JSON 并返回
	json.NewEncoder(w).Encode(article)
}

// 新增 getArticleList 处理函数
func getArticleList(w http.ResponseWriter, r *http.Request) {
	// 创建包含5条模拟数据的文章列表
	articles := []Article{
		{
			ID:      "1",
			Title:   "第一篇文章",
			Content: "这是第一篇文章的内容",
		},
		{
			ID:      "2",
			Title:   "第二篇文章",
			Content: "这是第二篇文章的内容",
		},
		{
			ID:      "3",
			Title:   "第三篇文章",
			Content: "这是第三篇文章的内容",
		},
		{
			ID:      "4",
			Title:   "第四篇文章",
			Content: "这是第四篇文章的内容",
		},
		{
			ID:      "5",
			Title:   "第五篇文章",
			Content: "这是第五篇文章的内容",
		},
	}

	// 设置响应头为 JSON
	w.Header().Set("Content-Type", "application/json")

	// 将文章列表转换为 JSON 并返回
	json.NewEncoder(w).Encode(articles)
}
