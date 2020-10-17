package mysqlrepo

import (
	"errors"
	"zampapp/lib/entity/model"
)

func (s Service) GetUser(idUser int) (model.User, error) {
	var u model.User
	res := s.gormDB.Find(&u, idUser)
	if res.RowsAffected != 1 {
		return u, errors.New("not found")
	}
	return u, res.Error
}

func (s Service) GetUserByLogin(username string) (model.User, error) {
	var u model.User
	tx := s.gormDB.Where("nick_name = ? OR email = ?", username, username).First(&u)
	return u, tx.Error
}
