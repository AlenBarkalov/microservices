package main
import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"github.com/gin-gonic/gin"
	)

var tmpArticleList []article

// This function is used for setup before executing the test functions
// Функция TestMain переводит Gin в тестовый режим и вызывает функции тестирования.
func TestMain(m *testing.M) {
	//Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Run the other tests
	os.Exit(m.Run())
}

// Helper function to create a router during testing
// Функция getRouter создаёт и возвращает роутер.
func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("templates/*")
	}
	return r
}

// Helper function to process a request and test its response
// функция testHTTPResponse выполняет переданную ей функцию для проверки — возвращает ли она
// булево значение true — показывая успешность теста, или нет.
// Эта функция помогает нам избежать дублирования кода, необходимого для тестирования ответа на HTTP-запрос.
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

// This function is used to store the main lists into the temporary one
// for testing
// Функция saveLists() помещает список топиков во временную переменную.
// Она используется в функции restoreLists() для восстановления списка топиков до первоначального состояния после выполнения юнит-теста.
func saveLists() {
	tmpArticleList = articleList
}

// This function is used to restore the main lists from the temporary one
func restoreLists() {
	articleList = tmpArticleList
}