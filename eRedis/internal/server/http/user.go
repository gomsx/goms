package http

import (
	"log"
	"strconv"

	"github.com/fuwensun/goms/eRedis/internal/model"
	"github.com/gin-gonic/gin"
)

func createUser(c *gin.Context) {
	var err error
	user := model.User{}

	namestr := c.Query("name")
	sexstr := c.Query("sex")

	user.Sex, err = strconv.ParseInt(sexstr, 10, 64)
	if sexstr == "" || err != nil {
		log.Printf("http sex err:%v\n", sexstr)
		c.JSON(404, gin.H{"error": "uid err!"})
		return
	}
	user.Name = namestr

	if err = svc.CreateUser(c, &user); err != nil {
		log.Printf("http create user, err: %v", err)
		c.JSON(404, gin.H{"error": "create failed!"})
		return
	}
	c.JSON(200, gin.H{
		"uid":  user.Uid,
		"name": user.Name,
		"sex":  user.Sex,
	})
	log.Printf("http create user=%v\n", user)
}

func updateUser(c *gin.Context) {
	var err error
	user := model.User{}

	uidstr := c.Query("uid")
	namestr := c.Query("name")
	sexstr := c.Query("sex")

	user.Uid, err = strconv.ParseInt(uidstr, 10, 64)
	if uidstr == "" || err != nil {
		log.Printf("http uid err:%v\n", uidstr)
		c.JSON(404, gin.H{"error": "uid err!"})
		return
	}
	user.Sex, err = strconv.ParseInt(sexstr, 10, 64)
	if sexstr == "" || err != nil {
		log.Printf("http sex err:%v\n", sexstr)
		c.JSON(404, gin.H{"error": "sex err!"})
		return
	}
	user.Name = namestr
	err = svc.UpdateUser(c, &user)
	if err != nil {
		log.Printf("http update user,err: %v\n", err)
		c.JSON(404, gin.H{"error": "data not found!"})
		return
	}
	c.JSON(200, gin.H{
		"uid":  user.Uid,
		"name": user.Name,
		"sex":  user.Sex,
	})
	log.Printf("http update user=%v\n", user)
}

func readUser(c *gin.Context) {
	uidstr := c.Query("uid")
	uid, err := strconv.ParseInt(uidstr, 10, 64)
	if uidstr == "" || err != nil {
		log.Printf("http uid err:%v\n", uidstr)
		c.JSON(404, gin.H{"error": "uid err!"})
		return
	}
	user, err := svc.ReadUser(c, uid)
	if err != nil {
		log.Printf("http read user,err: %v\n", err)
		c.JSON(404, gin.H{"error": "data not found!"})
		return
	}
	c.JSON(200, gin.H{
		"uid":  user.Uid,
		"name": user.Name,
		"sex":  user.Sex,
	})
	log.Printf("http read user=%v\n", user)
}

func deleteUser(c *gin.Context) {
	uidstr := c.Query("uid")
	uid, err := strconv.ParseInt(uidstr, 10, 64)
	if uidstr == "" || err != nil {
		log.Printf("http uid err:%v\n", uidstr)
		c.JSON(404, gin.H{"error": "uid err!"})
		return
	}
	if err = svc.DeleteUser(c, uid); err != nil {
		log.Printf("http delete user,err: %v\n", err)
		c.JSON(404, gin.H{"error": "data not found!"})
		return
	}
	c.JSON(200, gin.H{
		"uid": uid,
	})
	log.Printf("http delete user uid=%v\n", uid)
}
