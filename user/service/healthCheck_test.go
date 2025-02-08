package service_test

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	repo "poc.com/user/repository/mocks"
	"poc.com/user/service"
)

func MockCtx() *fiber.Ctx {
	return &fiber.Ctx{}
}

func TestHealthCheck(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := repo.NewMockUserRepository(ctrl)

	service := service.NewUserService(repoMock)

	t.Run("success", func(t *testing.T) {
		repoMock.EXPECT().GetAll().Return(nil)

		resp, err := service.HealthCheck(MockCtx())
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})
}
