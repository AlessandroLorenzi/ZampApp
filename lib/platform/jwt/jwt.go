package jwt

import (
	"zampapp/lib/entity/model"

	"github.com/brianvoe/sjwt"
)

var secretKey = "test"

func GenerateUserJWT(u model.User) string {
	claims := sjwt.New()
	claims.Set("nick_name", u.NickName)
	claims.Set("account_id", u.ID)

	// Generate jwt
	secretKey := []byte(secretKey)
	return claims.Generate(secretKey)
}
