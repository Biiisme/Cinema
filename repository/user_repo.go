package repository

import (
	"cinema/model"
	"cinema/model/req"
	"context"
)

type UserRepo interface {
	CheckLogin(context context.Context, loginReq req.ReqSignIn) (model.User, error)
	SaveUser(context context.Context, user model.User) (model.User, error)
	SaveToken(context context.Context, UserID string, token string, ip string) error
	GetUser(context context.Context, UserID string) (model.User, error)
	UpdateUser(filmReq req.ReqUpdateProfile, id string) (model.User, error)
	TotalPage(user model.User, length int) int
	GetAllUser(ctx context.Context, offset int, length int) ([]model.User, error)
}
