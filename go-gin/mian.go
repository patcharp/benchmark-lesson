package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
  gin.SetMode(gin.ReleaseMode)
  r := gin.Default()

  r.GET("/", func(c *gin.Context) {
	c.String(http.StatusOK, "go-gin")
  })

  r.Run(":3000")
}
