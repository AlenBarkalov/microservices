/*
Мы будет создавать все обработчики роутов, относящихся к топикам, в файле handlers.article.go.
Обработчик главной страницы, showIndexPage выполняет следующие задачи:
1. Получает список топиков
2. Обрабатывает шаблон index.html, передавая ему список топиков

 */

package main

import (
	"net/http"
	"strconv"
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

func getArticle(c *gin.Context) {
	// Проверим валидность ID
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Проверим существование топика
		if article, err := getArticleByID(articleID); err == nil {
			// Вызовем метод HTML из Контекста Gin для обработки шаблона
			c.HTML(
				// Зададим HTTP статус 200 (OK)
				http.StatusOK,
				// Используем шаблон index.html
				"article.html",
				// Передадим данные в шаблон
				gin.H{
					"title": article.Title,
					"payload": article,
				},
			)
		} else {
			// Если топика нет, прервём с ошибкой
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		// При некорректном ID в URL, прервём с ошибкой
		c.AbortWithStatus(http.StatusNotFound)
	}
}