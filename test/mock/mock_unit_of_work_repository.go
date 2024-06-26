// Code generated by MockGen. DO NOT EDIT.
// Source: ../internal/core/application/repository/unit_of_work_repository.go
//
// Generated by this command:
//
//	mockgen -source=../internal/core/application/repository/unit_of_work_repository.go -destination=../test/mock/mock_unit_of_work_repository.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockUnitOfWorkRepository is a mock of UnitOfWorkRepository interface.
type MockUnitOfWorkRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUnitOfWorkRepositoryMockRecorder
}

// MockUnitOfWorkRepositoryMockRecorder is the mock recorder for MockUnitOfWorkRepository.
type MockUnitOfWorkRepositoryMockRecorder struct {
	mock *MockUnitOfWorkRepository
}

// NewMockUnitOfWorkRepository creates a new mock instance.
func NewMockUnitOfWorkRepository(ctrl *gomock.Controller) *MockUnitOfWorkRepository {
	mock := &MockUnitOfWorkRepository{ctrl: ctrl}
	mock.recorder = &MockUnitOfWorkRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnitOfWorkRepository) EXPECT() *MockUnitOfWorkRepositoryMockRecorder {
	return m.recorder
}

// Begin mocks base method.
func (m *MockUnitOfWorkRepository) Begin(ctx context.Context) (context.Context, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Begin", ctx)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Begin indicates an expected call of Begin.
func (mr *MockUnitOfWorkRepositoryMockRecorder) Begin(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockUnitOfWorkRepository)(nil).Begin), ctx)
}

// Commit mocks base method.
func (m *MockUnitOfWorkRepository) Commit(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockUnitOfWorkRepositoryMockRecorder) Commit(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockUnitOfWorkRepository)(nil).Commit), ctx)
}

// Rollback mocks base method.
func (m *MockUnitOfWorkRepository) Rollback(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Rollback indicates an expected call of Rollback.
func (mr *MockUnitOfWorkRepositoryMockRecorder) Rollback(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockUnitOfWorkRepository)(nil).Rollback), ctx)
}
