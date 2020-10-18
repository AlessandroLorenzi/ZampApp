package mysqlrepo

import (
	"errors"
	"zampapp/lib/entity/model"
)

func (s Service) SaveUser(u model.User) error {
	tx := s.gormDB.Save(u)
	return tx.Error
}

func (s Service) GetUsers() (uu []model.User, err error) {
	tx := s.gormDB.Find(&uu)
	return uu, tx.Error
}

func (s Service) GetUser(idUser string) (model.User, error) {
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
