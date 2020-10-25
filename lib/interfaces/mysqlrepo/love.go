package mysqlrepo

import (
	"zampapp/lib/entity/model"

	"github.com/sirupsen/logrus"
)

func (s Service) CreateLove(l model.Love) error {
	tx := s.gormDB.Create(l)
	return tx.Error
}

func (s Service) RemoveLove(l model.Love) error {
	tx := s.gormDB.Delete(l)
	return tx.Error
}

func (s Service) AnimalsLovedBy(userId string) ([]model.Animal, error) {

	var (
		ll  []model.Love
		aa  []model.Animal
		a   model.Animal
		err error
	)

	tx := s.gormDB.Where("user = ?", userId).Find(&ll)
	if tx.Error != nil {
		return aa, tx.Error
	}

	for i := range ll {
		a, err = s.GetAnimal(ll[i].AnimalID)
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"error":  err.Error(),
				"func":   "animals loved by",
				"user":   userId,
				"animal": ll[i].AnimalID,
			}).Errorf("unexpected error")
			return []model.Animal{}, err
		}

		aa = append(aa, a)
	}

	return aa, nil
}
