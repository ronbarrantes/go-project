package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	// "ronb.co/project/database"
	"ronb.co/project/utils"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func run() {
	utils.LoadEnvs()
	// database.InitDB()

	PORT := os.Getenv("PORT")

	r := gin.Default()
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"hello": "world",
		})
	})

	log.Fatal(r.Run(":" + PORT))

}

func main() {
	run()
}
