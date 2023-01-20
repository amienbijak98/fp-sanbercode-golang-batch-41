package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wisedevguy/fp-sanbercode-golang-batch-41/database"
	"github.com/wisedevguy/fp-sanbercode-golang-batch-41/repository"
	"github.com/wisedevguy/fp-sanbercode-golang-batch-41/structs"
)

func GetAllItem(c *gin.Context) {
	var (
		result gin.H
	)

	items, err := repository.GetAllItem(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": items,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertItem(c *gin.Context) {
	var item structs.Item

	err := c.ShouldBindJSON(&item)
	if err != nil {
		panic(err)
	}

	err = repository.InsertItem(database.DbConnection, item)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success insert new item",
	})
}

func UpdateItem(c *gin.Context) {
	var item structs.Item
	id, _ := strconv.Atoi(c.Param("id_item"))

	err := c.ShouldBindJSON(&item)
	if err != nil {
		panic(err)
	}

	item.IDItem = int64(id)

	err = repository.UpdateItem(database.DbConnection, item)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success update item",
	})
}

func DeleteItem(c *gin.Context) {
	var item structs.Item
	idItem, _ := strconv.Atoi(c.Param("id_item"))

	item.IDItem = int64(idItem)

	err := repository.DeleteItem(database.DbConnection, item)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success delete item",
	})
}
