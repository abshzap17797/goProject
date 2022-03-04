package app
import (
	"github.com/gin-gonic/gin"
	"./domain"
)

var (
	router = gin.Default()
)

func StartApp() {
	MapUrls()
	domain.ConnDB()
	router.Run(":8080")
}
