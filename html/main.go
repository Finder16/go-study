package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Gin 엔진 생성
	router := gin.Default()

	// 정적 파일 서빙 설정
	router.Static("/static", "./static")

	// HTML 렌더링을 위한 템플릿 경로 설정
	router.LoadHTMLGlob("templates/*")

	// 루트 경로에 대한 핸들러
	router.GET("/", func(c *gin.Context) {
		// HTML 파일 렌더링
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// 서버 시작
	router.Run(":1")
}
