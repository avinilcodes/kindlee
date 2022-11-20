package state

import (
	"context"
	"kindlee/app"
	"kindlee/db"
	"strconv"

	"go.uber.org/zap"
)

type Service interface {
	getState(ctx context.Context, input StateRequest) (state StateResponse, err error)
}

type stateService struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

func checkState(rating int, inputRating int) int {
	if inputRating == 5 && rating <= 10 {
		return 1
	}
	if inputRating == 4 && rating <= 20 {
		return 1
	}
	if inputRating == 3 && rating <= 30 {
		return 1
	}
	if inputRating == 2 && rating <= 40 {
		return 1
	}
	if inputRating == 1 && rating <= 50 {
		return 1
	}
	return 0
}

func (cs *stateService) getState(ctx context.Context, input StateRequest) (state StateResponse, err error) {
	stateStats, err := cs.store.ListStats(ctx)
	if err != nil {
		app.GetLogger().Warn("Error while adding user", err.Error())
		return
	}

	var ss []db.State
	for _, j := range stateStats {
		jm, err := strconv.ParseFloat(j.MonthlyRentSpare, 64)
		if err != nil {
			state.State = ""
			return state, err
		}

		if jm < input.MonthlyRentSpare {
			ss = append(ss, j)
		}
	}
	stateStats = ss
	ss = ss[:0]
	for _, j := range stateStats {
		jf, err := strconv.Atoi(j.Financial)
		if err != nil {
			state.State = ""
			return state, err
		}
		jh, err := strconv.Atoi(j.Healthcare)
		if err != nil {
			state.State = ""
			return state, err
		}
		jp, err := strconv.Atoi(j.PersonalResidential)
		if err != nil {
			state.State = ""
			return state, err
		}
		jhc, err := strconv.Atoi(j.Healthcare)
		if err != nil {
			state.State = ""
			return state, err
		}
		jw, err := strconv.Atoi(j.Workplace)
		if err != nil {
			state.State = ""
			return state, err
		}

		sum := (checkState(jf, input.Financial) + checkState(jh, input.Healthcare) + checkState(jp, input.PersonalResidential) + checkState(jhc, input.Healthcare) + checkState(jw, input.Workplace))
		if sum > 2 {
			ss = append(ss, j)
		}
	}
	var s string
	for _, i := range ss {
		s += i.State + ","
	}
	state.State = s
	return state, nil
}
func NewService(s db.Storer, l *zap.SugaredLogger) Service {
	return &stateService{
		store:  s,
		logger: l,
	}
}
