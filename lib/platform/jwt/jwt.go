package jwt

import (
	"time"
	"zampapp/lib/entity/model"

	"github.com/brianvoe/sjwt"
)

var secretPrivateKey = []byte("test")
var secretPublicKey = []byte("test")

var duration = 6 * 24 * time.Hour

func GenerateUserJWT(u model.User) string {
	claims := sjwt.New()
	claims.Set("account_id", u.ID)
	claims.SetNotBeforeAt(time.Now().Add(duration))

	// Generate jwt
	return claims.Generate(secretPrivateKey)
}

func VerifyAndValidateUser(jwt string) string {
	if !sjwt.Verify(jwt, secretPublicKey) {
		return ""
	}

	claims, _ := sjwt.Parse(jwt)

	if claims.Validate() != nil {
		return ""
	}

	if idUser, err := claims.Get("account_id"); err != nil {
		return idUser.(string)
	}

	return ""
}
