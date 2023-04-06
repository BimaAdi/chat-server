package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func chatRoutes(rg *gin.RouterGroup) {
	chats := rg.Group("/chat")

	chats.GET("/", chatTemplate)
}

func chatTemplate(c *gin.Context) {
	c.HTML(http.StatusOK, "chat.html", nil)
}
