package main

// import "github.com/gin-gonic/gin"

// func Ping(c *gin.Context) { //c context
// 	number := c.Query("number")
// 	c.JSON(200, gin.H{"message": "Pong" + number})
// }

// func Hello(c *gin.Context) {
// 	c.JSON(200, gin.H{"message": "Hai"})
// }

// func Helo(c *gin.Context) {
// 	name := c.Param("name")
// 	c.JSON(200, gin.H{"message": "Hai" + name})
// }

// func main() { // r route
// 	r := gin.Default()

// 	r.GET("/ping", Ping)

// 	r.GET("/name/:name", Helo)

// 	r.GET("/hello", Hello)

// 	r.Run(":8080")
// }

import (
	"fmt"
	"go_gin_test/database"
	"go_gin_test/router"
	"log"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Gagal terhubung ke database: ", err)
	}
	defer db.Close()
	fmt.Println("✅ Terhubung ke database PostgreSQL")

	r := router.SetupRouter(db)

	fmt.Println("🚀 Server Gin berjalan di http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Gagal menjalankan server: ", err)
	}
}
