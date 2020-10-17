package usecases

import (
	"fmt"
	"zampapp/lib/entity/model"
)

func (s Service) Login(login, password string) (u model.User, err error) {
	u, err = s.repoService.GetUserByLogin(login)
	if err != nil {
		return
	}

	if u.ValidatePassword(password) {
		return u, nil
	}

	return model.User{}, fmt.Errorf("user and password don't match")
}
