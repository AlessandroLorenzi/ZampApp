package web

import (
	"github.com/gin-gonic/gin"
)

func (s *Service) GetUser(c *gin.Context) {
	idUser := c.Param("id_user")

	u, err := s.repoService.GetUser(idUser)
	if err != nil {
		if err.Error() == "not found" {
			c.JSON(400, gin.H{"msg": "user does not exists"})
			return
		}
		c.JSON(500, gin.H{"msg": "unexpected error"})
		return
	}

	c.JSON(200, gin.H{"msg": "ok", "user": u})
}
