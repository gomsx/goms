package http

import (
	"github.com/gin-gonic/gin"
)

func updatename(c *gin.Context) {
	// message := "pong" + " " + c.DefaultQuery("message", "NONE!")
	// c.JSON(200, gin.H{
	// 	"message": message,
	// })
	// log.Printf("http" + " " + message)
	svc.UpdateUserName(c, 1, "xxx")
}
