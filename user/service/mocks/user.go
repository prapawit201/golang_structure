// Code generated by MockGen. DO NOT EDIT.
// Source: user/service/init.go
//
// Generated by this command:
//
//	mockgen -source=user/service/init.go -destination=user/service/mocks/user.go
//

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	fiber "github.com/gofiber/fiber/v2"
	gomock "go.uber.org/mock/gomock"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
	isgomock struct{}
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// HealthCheck mocks base method.
func (m *MockUserService) HealthCheck(c *fiber.Ctx) (any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthCheck", c)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HealthCheck indicates an expected call of HealthCheck.
func (mr *MockUserServiceMockRecorder) HealthCheck(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockUserService)(nil).HealthCheck), c)
}
