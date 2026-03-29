package user

import (
	"sync"

	"ThisProject/models/repositories"
	userDataModel "ThisProject/models/user/dataModel"
)

type Repository struct {
	lock      sync.RWMutex
	idCounter int64
	users     []userDataModel.User
}

var (
	once    sync.Once
	repoIns *Repository
)

func GetRepo() *Repository {
	once.Do(func() {
		repoIns = &Repository{
			idCounter: 10,
			users:     []userDataModel.User{},
		}
	})

	return repoIns
}

func init() {
	repositories.UserRepo = GetRepo()
}
