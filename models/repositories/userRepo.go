package repositories

import (
	"context"

	"ThisProject/apiSchema/commonSchema"
	"ThisProject/apiSchema/userSchema"
)

type UserRepository interface {
	Create(ctx context.Context, req commonSchema.BaseRequest[userSchema.CreateRequest]) (res userSchema.CreateResponse, errStr string, code int, err error)
	Info(ctx context.Context, req commonSchema.BaseRequest[userSchema.InfoRequest]) (res userSchema.InfoResponse, errStr string, code int, err error)
}

var UserRepo UserRepository
