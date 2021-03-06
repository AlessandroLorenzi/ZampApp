package mysqlrepo

import (
	"errors"
	"zampapp/lib/entity/model"
)

func (s Service) CreateUser(u model.User) error {
	_, err := s.GetUserByLogin(u.Email)

	if err != nil && err.Error() != "record not found" {
		return err
	}
	if err == nil {
		return errors.New("email in use")
	}

	_, err = s.GetUserByLogin(u.NickName)
	if err != nil && err.Error() != "record not found" {
		return err
	}
	if err == nil {
		return errors.New("username in use")
	}

	tx := s.gormDB.Create(u)
	return tx.Error
}

func (s Service) GetUsers() (uu []model.User, err error) {
	tx := s.gormDB.Find(&uu)
	return uu, tx.Error
}

func (s Service) GetUser(idUser string) (model.User, error) {
	var u model.User
	res := s.gormDB.Where(`id = ?`, idUser).Find(&u)
	if res.RowsAffected != 1 {
		return u, errors.New("not found")
	}
	return u, res.Error
}

func (s Service) GetUserByLogin(nickOrEmail string) (model.User, error) {
	var u model.User

	tx := s.gormDB.Where("nick_name = ? OR email = ?", nickOrEmail, nickOrEmail).First(&u)
	if tx.Error != nil {
		return u, tx.Error
	}

	return u, tx.Error
}

func (s Service) DeleteUser(idUser string) error {
	tx := s.gormDB.Where("id = ?", idUser).Delete(model.User{})
	return tx.Error
}

func (s Service) UpdateUser(u model.User) error {
	u1, err := s.GetUserByLogin(u.Email)

	if err != nil && err.Error() != "record not found" {
		return err
	}
	if err == nil && u1.ID != u.ID {
		return errors.New("email in use")
	}

	u1, err = s.GetUserByLogin(u.NickName)
	if err != nil && err.Error() != "record not found" {
		return err
	}
	if err == nil && u1.ID != u.ID {
		return errors.New("username in use")
	}

	tx := s.gormDB.Save(u)
	return tx.Error
}
