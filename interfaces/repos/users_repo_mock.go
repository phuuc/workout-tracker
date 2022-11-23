// Code generated by MockGen. DO NOT EDIT.
// Source: users_repo.go

// Package repos is a generated GoMock package.
package repos

import (
	context "context"
	reflect "reflect"

	in "github.com/finnpn/workout-tracker/usecases/in"
	gomock "github.com/golang/mock/gomock"
)

// MockUsersRepo is a mock of UsersRepo interface.
type MockUsersRepo struct {
	ctrl     *gomock.Controller
	recorder *MockUsersRepoMockRecorder
}

// MockUsersRepoMockRecorder is the mock recorder for MockUsersRepo.
type MockUsersRepoMockRecorder struct {
	mock *MockUsersRepo
}

// NewMockUsersRepo creates a new mock instance.
func NewMockUsersRepo(ctrl *gomock.Controller) *MockUsersRepo {
	mock := &MockUsersRepo{ctrl: ctrl}
	mock.recorder = &MockUsersRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsersRepo) EXPECT() *MockUsersRepoMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockUsersRepo) Login(ctx context.Context, record *in.Login) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, record)
	ret0, _ := ret[0].(error)
	return ret0
}

// Login indicates an expected call of Login.
func (mr *MockUsersRepoMockRecorder) Login(ctx, record interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUsersRepo)(nil).Login), ctx, record)
}

// Register mocks base method.
func (m *MockUsersRepo) Register(ctx context.Context, record *in.Register) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, record)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockUsersRepoMockRecorder) Register(ctx, record interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockUsersRepo)(nil).Register), ctx, record)
}
