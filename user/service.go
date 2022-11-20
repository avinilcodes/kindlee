package user

import (
	"context"
	"kindlee/app"
	"kindlee/db"

	"go.uber.org/zap"
)

type Service interface {
	addUser(ctx context.Context, user db.User) (err error)
}

type userService struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

func (cs *userService) addUser(ctx context.Context, user db.User) (err error) {
	err = cs.store.CreateUser(ctx, user)
	if err != nil {
		app.GetLogger().Warn("Error while adding user", err.Error())
		return
	}
	return
}
func NewService(s db.Storer, l *zap.SugaredLogger) Service {
	return &userService{
		store:  s,
		logger: l,
	}
}
