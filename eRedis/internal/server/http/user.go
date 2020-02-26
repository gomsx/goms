package http

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func updatename(c *gin.Context) {

	uidstr := c.Query("uid")
	uid, err := strconv.ParseInt(uidstr, 10, 64)
	if uidstr == "" || err != nil {
		log.Printf("uid err:%v\n", uidstr)
		return
	}
	name := c.Query("name")
	if name == "" {
		log.Printf("name err:%v\n", uidstr)
		return
	}
	c.JSON(200, gin.H{
		"uid":  uidstr,
		"name": name,
	})
	log.Printf("http user updatename %v to %v\n", uidstr, name)
	svc.UpdateUserName(c, uid, name)
}

func readname(c *gin.Context) {

	uidstr := c.Query("uid")
	uid, err := strconv.ParseInt(uidstr, 10, 64)
	if uidstr == "" || err != nil {
		log.Printf("uid err:%v\n", uidstr)
		return
	}
	name, err := svc.ReadUserName(c, uid)
	if err != nil {
		c.JSON(404, gin.H{})
		return
	}
	c.JSON(200, gin.H{
		"uid":  uidstr,
		"name": name,
	})
	log.Printf("http user readname %v to %v\n", uidstr, name)
}
