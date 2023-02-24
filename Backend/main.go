package main

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"net/http"
	"v2/services"
)

func getvm(c *gin.Context) {
	services.getlist()
	c.IndentedJSON(http.StatusOK, services.vmlists)
	vmlists = vmlists[:0]
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/api/vm", services.getvm())

	router.Run("localhost:8080")

}
