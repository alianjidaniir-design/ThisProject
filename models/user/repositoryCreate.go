package user

import (
	"context"
	"sync/atomic"

	"ThisProject/apiSchema/commonSchema"
	"ThisProject/apiSchema/userSchema"
	userDataModel "ThisProject/models/user/dataModel"
	"ThisProject/statics/constants/status"
)

func (repo *Repository) nextID() int64 {
	return atomic.AddInt64(&repo.idCounter, 1)
}

func (repo *Repository) Create(ctx context.Context, req commonSchema.BaseRequest[userSchema.CreateRequest]) (res userSchema.CreateResponse, errStr string, code int, err error) {
	_ = ctx

	newUser := userDataModel.User{
		ID:       repo.nextID(),
		Username: req.Body.Username,
		Email:    req.Body.Email,
	}

	repo.lock.Lock()
	repo.users = append(repo.users, newUser)
	repo.lock.Unlock()

	return userSchema.CreateResponse{User: newUser}, "", status.StatusOK, nil
}
