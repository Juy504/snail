package main

import "snail/router"

func main() {
	/*r := gin.Default()
	// Heartbeat test
	r.GET("ping", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()*/

	r := router.SetupRouter()

	// running
	r.Run()
}
