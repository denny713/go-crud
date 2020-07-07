package controller

import (
	"../util"
	"Github/go-crud/model"
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
	user.Password = util.EncryptText(user.Password)
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
	userUpd.Password = util.EncryptText(user.Password)
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
