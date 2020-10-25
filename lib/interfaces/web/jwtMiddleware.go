package web

import (
	"time"
	"zampapp/lib/entity/model"

	jwt "github.com/appleboy/gin-jwt/v2"

	"github.com/gin-gonic/gin"
)

var identityKey = "user_id"

func (s Service) jwtMiddleware() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:         "test zone",
		Key:           []byte("secret key"),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		IdentityKey:   identityKey,
		PayloadFunc:   s.jwtPayload,
		Authenticator: s.login,

		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}
}

func (s *Service) jwtPayload(data interface{}) jwt.MapClaims {
	if v, ok := data.(*model.User); ok {
		return jwt.MapClaims{
			identityKey: v.ID,
		}
	}
	return jwt.MapClaims{}
}

func (s *Service) login(c *gin.Context) (interface{}, error) {
	type loginPost struct {
		Login    string `json:"login" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var lp loginPost
	if err := c.ShouldBind(&lp); err != nil {
		return "", jwt.ErrMissingLoginValues
	}

	u, err := s.useCasesService.Login(lp.Login, lp.Password)
	if err != nil {
		return u, err
	}
	return nil, jwt.ErrFailedAuthentication
}