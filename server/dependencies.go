package server

import (
	"kindlee/app"
	"kindlee/db"
	"kindlee/login"
	"kindlee/user"
	"kindlee/utils"
)

type dependencies struct {
	UserLoginService login.Service
	UserServices     user.Service
}

func initDependencies() (dependencies, error) {
	appDB := app.GetDB()
	logger := app.GetLogger()
	dbStore := db.NewStorer(appDB)

	// call new service
	userService := user.NewService(dbStore, logger)
	loginService := login.NewService(dbStore, logger)

	err := db.CreateSuperAdmin(dbStore)
	if err != nil && !utils.CheckIfDuplicateKeyError(err) {
		return dependencies{}, err
	}

	return dependencies{
		UserLoginService: loginService,
		UserServices:     userService,
	}, nil
}
