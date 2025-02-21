// Code generated by MockGen. DO NOT EDIT.
// Source: IUser.go

// Package services is a generated GoMock package.
package services

import (
	context "context"
	models "ewallet-ums/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIUserRepository is a mock of IUserRepository interface.
type MockIUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIUserRepositoryMockRecorder
}

// MockIUserRepositoryMockRecorder is the mock recorder for MockIUserRepository.
type MockIUserRepositoryMockRecorder struct {
	mock *MockIUserRepository
}

// NewMockIUserRepository creates a new mock instance.
func NewMockIUserRepository(ctrl *gomock.Controller) *MockIUserRepository {
	mock := &MockIUserRepository{ctrl: ctrl}
	mock.recorder = &MockIUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserRepository) EXPECT() *MockIUserRepositoryMockRecorder {
	return m.recorder
}

// DeleteUserSession mocks base method.
func (m *MockIUserRepository) DeleteUserSession(ctx context.Context, token string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserSession", ctx, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserSession indicates an expected call of DeleteUserSession.
func (mr *MockIUserRepositoryMockRecorder) DeleteUserSession(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserSession", reflect.TypeOf((*MockIUserRepository)(nil).DeleteUserSession), ctx, token)
}

// GetUserByUsername mocks base method.
func (m *MockIUserRepository) GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsername", ctx, username)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUsername indicates an expected call of GetUserByUsername.
func (mr *MockIUserRepositoryMockRecorder) GetUserByUsername(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsername", reflect.TypeOf((*MockIUserRepository)(nil).GetUserByUsername), ctx, username)
}

// GetUserSessionByRefreshToken mocks base method.
func (m *MockIUserRepository) GetUserSessionByRefreshToken(ctx context.Context, refreshToken string) (models.UserSession, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserSessionByRefreshToken", ctx, refreshToken)
	ret0, _ := ret[0].(models.UserSession)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserSessionByRefreshToken indicates an expected call of GetUserSessionByRefreshToken.
func (mr *MockIUserRepositoryMockRecorder) GetUserSessionByRefreshToken(ctx, refreshToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserSessionByRefreshToken", reflect.TypeOf((*MockIUserRepository)(nil).GetUserSessionByRefreshToken), ctx, refreshToken)
}

// GetUserSessionByToken mocks base method.
func (m *MockIUserRepository) GetUserSessionByToken(ctx context.Context, token string) (models.UserSession, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserSessionByToken", ctx, token)
	ret0, _ := ret[0].(models.UserSession)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserSessionByToken indicates an expected call of GetUserSessionByToken.
func (mr *MockIUserRepositoryMockRecorder) GetUserSessionByToken(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserSessionByToken", reflect.TypeOf((*MockIUserRepository)(nil).GetUserSessionByToken), ctx, token)
}

// InsertNewUser mocks base method.
func (m *MockIUserRepository) InsertNewUser(ctx context.Context, user *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertNewUser", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertNewUser indicates an expected call of InsertNewUser.
func (mr *MockIUserRepositoryMockRecorder) InsertNewUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertNewUser", reflect.TypeOf((*MockIUserRepository)(nil).InsertNewUser), ctx, user)
}

// InsertNewUserSession mocks base method.
func (m *MockIUserRepository) InsertNewUserSession(ctx context.Context, session *models.UserSession) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertNewUserSession", ctx, session)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertNewUserSession indicates an expected call of InsertNewUserSession.
func (mr *MockIUserRepositoryMockRecorder) InsertNewUserSession(ctx, session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertNewUserSession", reflect.TypeOf((*MockIUserRepository)(nil).InsertNewUserSession), ctx, session)
}

// UpdateTokenByRefreshToken mocks base method.
func (m *MockIUserRepository) UpdateTokenByRefreshToken(ctx context.Context, token, refreshToken string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTokenByRefreshToken", ctx, token, refreshToken)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTokenByRefreshToken indicates an expected call of UpdateTokenByRefreshToken.
func (mr *MockIUserRepositoryMockRecorder) UpdateTokenByRefreshToken(ctx, token, refreshToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTokenByRefreshToken", reflect.TypeOf((*MockIUserRepository)(nil).UpdateTokenByRefreshToken), ctx, token, refreshToken)
}
