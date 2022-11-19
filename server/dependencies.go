package server

import (
	"taskmanager/app"
	"taskmanager/db"
	"taskmanager/login"
	"taskmanager/utils"
)

type dependencies struct {
	UserLoginService login.Service
}

func initDependencies() (dependencies, error) {
	appDB := app.GetDB()
	logger := app.GetLogger()
	dbStore := db.NewStorer(appDB)

	// call new service

	loginService := login.NewService(dbStore, logger)

	err := db.CreateSuperAdmin(dbStore)
	if err != nil && !utils.CheckIfDuplicateKeyError(err) {
		return dependencies{}, err
	}

	return dependencies{
		UserLoginService: loginService,
	}, nil
}
