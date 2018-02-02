/*
	Сделаем структуру топика простой,
	всего с тремя полями — Id, Title (название) и Content (содержание).
 */
package main

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

/*
	Большинство приложений используют базу данных для хранения данных.
	Чтобы не усложнять, мы будем хранить список топиков в памяти и заполнять его при создании двумя следующими топиками:
 */

var articleList = []article{
	article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

// Return a list of all the articles
func getAllArticles() []article {
	return articleList
}

// Получить сам топик
func getArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}