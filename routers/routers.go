package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/wisedevguy/fp-sanbercode-golang-batch-41/controllers"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	// ROUTERS USERS
	router.GET("/users", controllers.GetAllUsers)
	router.GET("/users/:username", controllers.GetUserByUsername)
	router.POST("/users", controllers.InsertUser)
	router.PUT("/users/:username", controllers.UpdateUser)
	router.DELETE("/users/:username", controllers.DeleteUser)

	// ROUTERS ITEM
	router.GET("/items", controllers.GetAllItem)
	router.POST("/items", controllers.InsertItem)
	router.PUT("/items/:id_item", controllers.UpdateItem)
	router.DELETE("/items/:id_item", controllers.DeleteItem)

	return router
}
