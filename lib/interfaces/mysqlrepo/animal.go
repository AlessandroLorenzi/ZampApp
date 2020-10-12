package mysqlrepo

import (
	"errors"
	"zampapp/lib/entity/model"
)

func (s Service) GetAnimal(idAnimal int) (model.Animal, error) {
	var a model.Animal
	res := s.gormDB.Find(&a, idAnimal)
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
