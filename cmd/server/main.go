//Application entry point - Every Go application starts from main().

package main

import (
	"fmt"

	"food-delivery-order/internal/config"

	"food-delivery-order/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	database.Connect(cfg)

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server Running",
		})
	})

	fmt.Println("Running on port", cfg.Port)

	router.Run(":" + cfg.Port)
}

/*
	1. Every Go file belongs to a package.
	2. main tells Go: "This is the program's entry point." Go automatically calls main() when your program starts.
	3. := Create a new variable called router, and let Go figure out its type. Go infers the type automatically.
	4. In Go with Gin, it's very similar, but req and res are combined into a single context object called c.
	5. gin.H is just a shortcut for creating a JSON object (a map)
*/
