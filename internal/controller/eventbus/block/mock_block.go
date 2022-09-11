// Code generated by MockGen. DO NOT EDIT.
// Source: block.go

// Package block is a generated GoMock package.
package block

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	metadata "github.com/linkall-labs/vanus/internal/controller/eventbus/metadata"
	kv "github.com/linkall-labs/vanus/internal/kv"
)

// MockAllocator is a mock of Allocator interface.
type MockAllocator struct {
	ctrl     *gomock.Controller
	recorder *MockAllocatorMockRecorder
}

// MockAllocatorMockRecorder is the mock recorder for MockAllocator.
type MockAllocatorMockRecorder struct {
	mock *MockAllocator
}

// NewMockAllocator creates a new mock instance.
func NewMockAllocator(ctrl *gomock.Controller) *MockAllocator {
	mock := &MockAllocator{ctrl: ctrl}
	mock.recorder = &MockAllocatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAllocator) EXPECT() *MockAllocatorMockRecorder {
	return m.recorder
}

// Pick mocks base method.
func (m *MockAllocator) Pick(ctx context.Context, num int) ([]*metadata.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pick", ctx, num)
	ret0, _ := ret[0].([]*metadata.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Pick indicates an expected call of Pick.
func (mr *MockAllocatorMockRecorder) Pick(ctx, num interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pick", reflect.TypeOf((*MockAllocator)(nil).Pick), ctx, num)
}

// Run mocks base method.
func (m *MockAllocator) Run(ctx context.Context, kvCli kv.Client, dynamicAllocate bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", ctx, kvCli, dynamicAllocate)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockAllocatorMockRecorder) Run(ctx, kvCli, dynamicAllocate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockAllocator)(nil).Run), ctx, kvCli, dynamicAllocate)
}

// Stop mocks base method.
func (m *MockAllocator) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockAllocatorMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockAllocator)(nil).Stop))
}