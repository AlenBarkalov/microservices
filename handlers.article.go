/*
Мы будет создавать все обработчики роутов, относящихся к топикам, в файле handlers.article.go.
Обработчик главной страницы, showIndexPage выполняет следующие задачи:
1. Получает список топиков
2. Обрабатывает шаблон index.html, передавая ему список топиков

 */

package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()						// 1. Получает список топиков

	// Call the HTML method of the Context to render a template
	// 2. Обрабатывает шаблон index.html, передавая ему список топиков
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,

		// Use the index.html template
		"index.html",

		// Pass the data that the page uses
		gin.H{
			"title": "Home Page",
			"payload": articles,
		},
	)
}