package main

import (
	"BinaryDecoder/src/decode"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type DecodeRequest struct {
	Binary string `json:"binary"`
	Data   string `json:"data"`
}

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	r.POST("/api/decode", func(c *gin.Context) {
		var req DecodeRequest

		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, "invalid request format")
			return
		}

		decoded, err := decode.DecodeHandler(req.Binary, req.Data)
		if err != nil {
			c.JSON(http.StatusBadRequest, "decode failed")
			return
		}

		c.JSON(http.StatusOK, decoded)
	})

	r.Run(":2500")
}
