package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wisedevguy/fp-sanbercode-golang-batch-41/database"
	"github.com/wisedevguy/fp-sanbercode-golang-batch-41/repository"
	"github.com/wisedevguy/fp-sanbercode-golang-batch-41/structs"
)

func GetAllUsers(c *gin.Context) {
	var (
		result gin.H
	)

	users, err := repository.GetAllUser(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": users,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertUser(c *gin.Context) {
	var user structs.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		panic(err)
	}

	err = repository.InsertUser(database.DbConnection, user)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success insert new user",
	})
}

func UpdateUser(c *gin.Context) {
	var user structs.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		panic(err)
	}

	user.Username = c.Param("username")

	err = repository.UpdateUser(database.DbConnection, user)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success update User",
	})
}

func DeleteUser(c *gin.Context) {
	var user structs.User
	user.Username = c.Param("username")

	err := repository.DeleteUser(database.DbConnection, user)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success delete user",
	})
}

func GetUserByUsername(c *gin.Context) {
	var (
		result gin.H
		user   structs.User
	)

	user.Username = c.Param("username")
	userOut, err := repository.GetUserByUsername(database.DbConnection, user)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": userOut,
		}
	}

	c.JSON(http.StatusOK, result)
}
