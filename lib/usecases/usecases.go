package usecases

import "zampapp/lib/entity/model"

type Service struct {
	repoService repo
}

type repo interface {
	GetUserByLogin(string) (model.User, error)
}

func New(repoService repo) Service {
	return Service{
		repoService: repoService,
	}
}
