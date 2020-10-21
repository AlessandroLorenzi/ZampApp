package usecases

import (
	"zampapp/lib/interfaces/mysqlrepo"
)

type Service struct {
	repoService mysqlrepo.Service
}

func New(repoService mysqlrepo.Service) Service {
	return Service{
		repoService: repoService,
	}
}
