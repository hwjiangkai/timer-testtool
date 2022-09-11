// Code generated by MockGen. DO NOT EDIT.
// Source: offset.go

// Package storage is a generated GoMock package.
package storage

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	info "github.com/linkall-labs/vanus/internal/primitive/info"
	vanus "github.com/linkall-labs/vanus/internal/primitive/vanus"
)

// MockOffsetStorage is a mock of OffsetStorage interface.
type MockOffsetStorage struct {
	ctrl     *gomock.Controller
	recorder *MockOffsetStorageMockRecorder
}

// MockOffsetStorageMockRecorder is the mock recorder for MockOffsetStorage.
type MockOffsetStorageMockRecorder struct {
	mock *MockOffsetStorage
}

// NewMockOffsetStorage creates a new mock instance.
func NewMockOffsetStorage(ctrl *gomock.Controller) *MockOffsetStorage {
	mock := &MockOffsetStorage{ctrl: ctrl}
	mock.recorder = &MockOffsetStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOffsetStorage) EXPECT() *MockOffsetStorageMockRecorder {
	return m.recorder
}

// CreateOffset mocks base method.
func (m *MockOffsetStorage) CreateOffset(ctx context.Context, subscriptionID vanus.ID, info info.OffsetInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOffset", ctx, subscriptionID, info)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOffset indicates an expected call of CreateOffset.
func (mr *MockOffsetStorageMockRecorder) CreateOffset(ctx, subscriptionID, info interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOffset", reflect.TypeOf((*MockOffsetStorage)(nil).CreateOffset), ctx, subscriptionID, info)
}

// DeleteOffset mocks base method.
func (m *MockOffsetStorage) DeleteOffset(ctx context.Context, subscriptionID vanus.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOffset", ctx, subscriptionID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOffset indicates an expected call of DeleteOffset.
func (mr *MockOffsetStorageMockRecorder) DeleteOffset(ctx, subscriptionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOffset", reflect.TypeOf((*MockOffsetStorage)(nil).DeleteOffset), ctx, subscriptionID)
}

// GetOffsets mocks base method.
func (m *MockOffsetStorage) GetOffsets(ctx context.Context, subscriptionID vanus.ID) (info.ListOffsetInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOffsets", ctx, subscriptionID)
	ret0, _ := ret[0].(info.ListOffsetInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOffsets indicates an expected call of GetOffsets.
func (mr *MockOffsetStorageMockRecorder) GetOffsets(ctx, subscriptionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOffsets", reflect.TypeOf((*MockOffsetStorage)(nil).GetOffsets), ctx, subscriptionID)
}

// UpdateOffset mocks base method.
func (m *MockOffsetStorage) UpdateOffset(ctx context.Context, subscriptionID vanus.ID, info info.OffsetInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOffset", ctx, subscriptionID, info)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOffset indicates an expected call of UpdateOffset.
func (mr *MockOffsetStorageMockRecorder) UpdateOffset(ctx, subscriptionID, info interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOffset", reflect.TypeOf((*MockOffsetStorage)(nil).UpdateOffset), ctx, subscriptionID, info)
}