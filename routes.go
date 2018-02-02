package main

//import "github.com/gin-gonic/gin"

//func initializeRoutes() *gin.Engine{
func initializeRoutes() {
	// определение роута главной страницы
	//r := gin.Default()
	router.GET("/", showIndexPage)
	//r.GET("/", showIndexPage)
	//return r
}