package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wisedevguy/fp-sanbercode-golang-batch-41/database"
	"github.com/wisedevguy/fp-sanbercode-golang-batch-41/repository"
	"github.com/wisedevguy/fp-sanbercode-golang-batch-41/structs"
)

func GetAllTransaksi(c *gin.Context) {
	var (
		result gin.H
	)

	trxs, err := repository.GetAllTransaksi(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": trxs,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertTransaksiCetak(c *gin.Context) {
	var transaksi structs.Transaksi

	err := c.ShouldBindJSON(&transaksi)
	if err != nil {
		panic(err)
	}

	err = repository.InsertTransaksiCetak(database.DbConnection, transaksi)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success insert new transaksi row",
	})
}
