package main

import (
	"context"
	"example51/schedule"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	s := schedule.New()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/cancel-task/:id", func(c *gin.Context) {
		taskID := c.Param("id")

		if err := s.Cancel(context.Background(), taskID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "pong",
			})
			return
		}

		// Update the task status.

		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	r.GET("/watch-task/:id", func(c *gin.Context) {
		taskID := c.Param("id")

		ctxDone, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		ok, _ := s.Cancelled(ctxDone, taskID)
		if ok {
			c.JSON(http.StatusOK, gin.H{
				"cancel": true,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"cancel": false,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}