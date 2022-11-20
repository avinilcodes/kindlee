package state

import (
	"context"
	"kindlee/app"
	"kindlee/db"

	"go.uber.org/zap"
)

type Service interface {
	getState(ctx context.Context, input StateRequest) (state StateResponse, err error)
}

type stateService struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

func (cs *stateService) getState(ctx context.Context, input StateRequest) (state StateResponse, err error) {
	stateStats, err := cs.store.ListStats(ctx)
	if err != nil {
		app.GetLogger().Warn("Error while adding user", err.Error())
		return
	}
	state.State = stateStats[0].State
	return state, nil
}
func NewService(s db.Storer, l *zap.SugaredLogger) Service {
	return &stateService{
		store:  s,
		logger: l,
	}
}
