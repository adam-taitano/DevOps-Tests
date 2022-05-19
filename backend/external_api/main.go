package main

import (
	"external_api/api"
	"external_api/db"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	pgdb := db.EnvConnect()
	log.SetFlags(2)
	router.Use(gin.LoggerWithFormatter(
		func(param gin.LogFormatterParams) string {
			return fmt.Sprintf("[XAPI INFO] %s %s -> %d\n", param.Method, param.Path, param.StatusCode)
		}))
	router.Use(gin.Recovery())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api.GroupApi(router, pgdb)

	router.Run(fmt.Sprintf(":%d", pgdb.Port))
	pgdb.Close()
}
