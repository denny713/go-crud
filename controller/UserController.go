package controller

import (
	"Github/go-crud/model"
	"crypto/sha256"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(c *gin.Context) {
	var users []model.User
	err := model.GetAllUser(&users)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func GetUserByUsername(c *gin.Context) {
	var user model.User
	id := c.Params.ByName("id")
	err := model.GetByUsername(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func SaveUser(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)
	pass := []byte(user.Password)
	hash := sha256.Sum256(pass)
	user.Password = fmt.Sprintf("%x", hash)
	err := model.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Success Insert User With Username " + user.Username: "",
		})
	}
}

func UpdateUser(c *gin.Context) {
	var user model.User
	var userUpd model.User
	id := c.Params.ByName("id")
	err := model.GetByUsername(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, "User Not Found")
	}
	c.BindJSON(&user)
	userUpd = user
	pass := []byte(user.Password)
	hash := sha256.Sum256(pass)
	userUpd.Password = fmt.Sprintf("%x", hash)
	err = model.DeleteUser(&user, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	err = model.CreateUser(&userUpd)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Success Update User With Username " + id: "",
		})
	}
}

func DeleteUser(c *gin.Context) {
	var user model.User
	id := c.Params.ByName("id")
	err := model.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Success Delete User With Username " + id: "",
		})
	}
}
