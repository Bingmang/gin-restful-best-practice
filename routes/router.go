package routes

import "github.com/gin-gonic/gin"

var engine = gin.Default()

func New() *gin.Engine {
	return engine
}
