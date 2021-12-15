package IndexController

import (
	"app/Services/UserService"
	"github.com/gin-gonic/gin"
	"os"
)

func Index(response *gin.Context) {
	response.JSON(200, gin.H{
		"version": os.Getenv("API_VERSION"),
		"MODE":    os.Getenv("GIN_MODE"),
		"message": "PING",
	})

	UserService.CreateTestUser()
}
