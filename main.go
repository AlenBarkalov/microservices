package main

import (
	"github.com/gin-gonic/gin"
	//"net/http"
)

var router *gin.Engine

func main() {
	// 1. Создаём роутер Gin
	router := gin.Default()

	// 2. Загрузим все шаблоны
	// Обработаем шаблоны вначале, так что их не нужно будет перечитывать
	// ещё раз. Из-за этого вывод HTML-страниц такой быстрый.
	router.LoadHTMLGlob("templates/*")
	/*
		Это загрузит все шаблоны из папки templates. Загрузив один раз шаблоны,
		больше не будет необходимости перечитывать их, что делает веб-приложения с Gin очень быстрыми.
	*/

	// 3. Задаём обработчик роутов:
	/* Выносим содержимое в routes.go
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
	*/

	//initializeRoutes()
	router.GET("/", showIndexPage)

	// 4. Запуск приложения
	router.Run(":8081")		// задаём отдельный порт
}
