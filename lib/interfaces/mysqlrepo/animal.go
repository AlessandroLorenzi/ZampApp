package mysqlrepo

import (
	"errors"
	"zampapp/lib/entity/model"
)

func (s Service) CreateAnimal(a model.Animal) error {
	tx := s.gormDB.Create(a)
	return tx.Error
}

func (s Service) GetAnimal(idAnimal string) (model.Animal, error) {
	var a model.Animal
	res := s.gormDB.Where("id = ?", idAnimal).Find(&a)
	if res.Error != nil {
		return a, res.Error
	}

	if res.RowsAffected != 1 {
		return a, errors.New("not found")
	}

	var err error
	a.Owner, err = s.GetUser(a.OwnerID)
	if err != nil {
		return a, err
	}

	return a, res.Error
}

func (s Service) GetAnimals() ([]model.Animal, error) {
	var aa []model.Animal
	res := s.gormDB.Find(&aa)
	for index := range aa {
		var err error
		aa[index].Owner, err = s.GetUser(aa[index].OwnerID)
		if err != nil {
			return aa, err
		}
	}
	return aa, res.Error
}

func (s Service) DeleteAnimal(idAnimal string) error {
	tx := s.gormDB.Where("id = ?", idAnimal).Delete(model.Animal{})
	return tx.Error
}

func (s Service) UpdateAnimal(a model.Animal) error {
	tx := s.gormDB.Save(a)
	return tx.Error
}
