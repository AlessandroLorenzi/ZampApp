package web

import (
	"zampapp/lib/platform/jwt"

	"github.com/gin-gonic/gin"
)

func (s *Service) login(c *gin.Context) {

	type loginPost struct {
		Login    string `json:"login" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var lp loginPost
	if err := c.ShouldBindJSON(&lp); err != nil {
		s.logger.Debug("Error login", err)
		c.JSON(400, gin.H{
			"msg": "login not valid",
		})
		return
	}

	u, err := s.useCasesService.Login(lp.Login, lp.Password)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "login not valid",
		})
		return
	}

	c.JSON(200, gin.H{"msg": "token generated", "content": gin.H{
		"user": u,
		"jwt":  jwt.GenerateUserJWT(u),
	}})

}
