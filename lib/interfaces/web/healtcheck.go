package web

import "github.com/gin-gonic/gin"

func (s Service) healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "ok"})
}
