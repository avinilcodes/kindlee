package user

import (
	"context"
	"kindlee/app"
	"kindlee/db"
	storemocks "kindlee/db/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type ServiceTestSuite struct {
	suite.Suite
	store   *storemocks.Storer
	logger  *zap.SugaredLogger
	service Service
}

func init() {
	app.InitLogger()
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

func (suite *ServiceTestSuite) SetupTest() {
	suite.logger = app.GetLogger()
	suite.store = &storemocks.Storer{}
	suite.service = NewService(suite.store, suite.logger)
}

func (suite *ServiceTestSuite) TearDownTest() {
	suite.store.AssertExpectations(suite.T())
}

func (suite *ServiceTestSuite) TestAddUser() {
	t := suite.T()
	var user db.User
	user.Name = "Sachin"
	user.Email = "Sachin@Josh.com"
	user.Password = "Sachin@Josh"
	user.RoleType = "super_admin"
	ctx := context.Background()

	suite.store.On("CreateUser", ctx, user).Return(nil)

	gotErr := suite.service.addUser(ctx, user)
	assert.NoError(t, gotErr)
}
