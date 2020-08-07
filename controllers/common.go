package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
)

type testResponse struct {
	*httptest.ResponseRecorder
	Data gin.H
}

func testRequest(controller func(*gin.Context), method, url string, params gin.H) *testResponse {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	if method == "GET" {
		query := "?"
		first := true
		for k, v := range params {
			if first {
				first = false
			} else {
				query += "&"
			}
			query += fmt.Sprintf("%s=%v", k, v)
		}
		c.Request, _ = http.NewRequest(method, url+query, nil)
	} else {
		jsonBytes, _ := json.Marshal(params)
		c.Request, _ = http.NewRequest(method, url, bytes.NewBuffer(jsonBytes))
	}
	c.Request.Header.Set("Content-Type", "application/json")
	controller(c)
	var data gin.H
	_ = json.Unmarshal(response.Body.Bytes(), &data)
	return &testResponse{
		ResponseRecorder: response,
		Data:             data,
	}
}

func ErrorResponse(ctx *gin.Context, statusCode int, message string) {
	ctx.JSON(statusCode, gin.H{
		"error": message,
	})
}
