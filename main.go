package main

import (
	"github.com/gin-gonic/gin"
	"github.com/suhailmshaik/pg-crn-test/config"
	"github.com/suhailmshaik/pg-crn-test/routes"
)

func main() {
	// Initializing a GIN router or a Web Server
	router:=gin.New()

	// Connecting to the database
	config.Connect()


	// // Passing router or web server into user defined package of routes
	routes.PayoutRouter(router)


	// Running the server on port 5000
	router.Run(":5000")
}