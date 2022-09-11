// Code generated by MockGen. DO NOT EDIT.
// Source: reader.go

// Package reader is a generated GoMock package.
package reader

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	info "github.com/linkall-labs/vanus/internal/primitive/info"
)

// MockReader is a mock of Reader interface.
type MockReader struct {
	ctrl     *gomock.Controller
	recorder *MockReaderMockRecorder
}

// MockReaderMockRecorder is the mock recorder for MockReader.
type MockReaderMockRecorder struct {
	mock *MockReader
}

// NewMockReader creates a new mock instance.
func NewMockReader(ctrl *gomock.Controller) *MockReader {
	mock := &MockReader{ctrl: ctrl}
	mock.recorder = &MockReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReader) EXPECT() *MockReaderMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockReader) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockReaderMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockReader)(nil).Close))
}

// GetOffsetByTimestamp mocks base method.
func (m *MockReader) GetOffsetByTimestamp(ctx context.Context, timestamp int64) (info.ListOffsetInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOffsetByTimestamp", ctx, timestamp)
	ret0, _ := ret[0].(info.ListOffsetInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOffsetByTimestamp indicates an expected call of GetOffsetByTimestamp.
func (mr *MockReaderMockRecorder) GetOffsetByTimestamp(ctx, timestamp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOffsetByTimestamp", reflect.TypeOf((*MockReader)(nil).GetOffsetByTimestamp), ctx, timestamp)
}

// Start mocks base method.
func (m *MockReader) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockReaderMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockReader)(nil).Start))
}