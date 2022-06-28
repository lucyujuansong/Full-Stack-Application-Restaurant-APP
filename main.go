package main

import (
	"os"
	"server1/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8081"
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.Use(cors.Default())

	// these are the endpoints
	// C
	router.POST("/order/create", routes.AddOrder)
	// R
	router.GET("/waiter/:waiter", routes.GetOrdersByWaiter)
	router.GET("/orders", routes.GetOrders)
	router.GET("/order/:id/", routes.GetOrderById)
	// U
	router.PUT("/waiter/upate/:id", routes.UpdateWaiter)
	router.PUT("/order/update/:id", routes.UpdateOrder)
	// D
	router.DELETE("/order/delete/:id", routes.DeleteOrder)

	// this runs the server and allows it to listen to requests
	router.Run(":" + port)
}
