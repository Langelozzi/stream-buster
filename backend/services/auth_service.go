package services

import (
	iDao "github.com/STREAM-BUSTER/stream-buster/daos/interfaces"
)

type AuthService struct {
	dao iDao.MovieDatabaseDaoInterface
}
