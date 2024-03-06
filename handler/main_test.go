package handler

// import (
// 	"fmt"
// 	database "homework/data_base"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestHello(t *testing.T) {
// 	Db := database.NewTempDataBase()
// 	handler := handler.NewHandler(Db)
// 	server := setupRoute(handler)

// 	req, _ := http.NewRequest("GET", "/tasks", nil)
// 	w := httptest.NewRecorder()
// 	server.ServeHTTP(w, req)

// 	expectedStatus := http.StatusOK
// 	assert.Equal(t, expectedStatus, w.Code)

// 	fmt.Printf("Body: %+v", w.Body)
// }
