package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var router *gin.Engine

func main() {
	// 1. Создаём роутер Gin
	router := gin.Default()

	// 2. Загрузим все шаблоны
	router.LoadHTMLGlob("templates/*")
	/*
		Это загрузит все шаблоны из папки templates. Загрузив один раз шаблоны,
		больше не будет необходимости перечитывать их, что делает веб-приложения с Gin очень быстрыми.
	*/

	// 3. Задаём обработчик роутов:
	router.GET("/", func(c *gin.Context) {

		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"index.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title": "Домашняя страница",
			},
		)

	})

	// 4. Запуск приложения
	router.Run(":8081")		// задаём отдельный порт
}
