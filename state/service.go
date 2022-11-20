package state

import (
	"context"
	"kindlee/app"
	"kindlee/db"

	"go.uber.org/zap"
)

type Service interface {
	getState(ctx context.Context, input StateRequest) (err error)
}

type stateService struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

func (cs *stateService) getState(ctx context.Context, input StateRequest) (err error) {
	stateStats, err := cs.store.ListStats(ctx)
	if err != nil {
		app.GetLogger().Warn("Error while adding user", err.Error())
		return
	}
	return
}
func NewService(s db.Storer, l *zap.SugaredLogger) Service {
	return &stateService{
		store:  s,
		logger: l,
	}
}
