// Code generated by MockGen. DO NOT EDIT.
// Source: connection.go

// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	protocol "github.com/hugefiver/quic/internal/protocol"
)

// MockStreamGetter is a mock of StreamGetter interface.
type MockStreamGetter struct {
	ctrl     *gomock.Controller
	recorder *MockStreamGetterMockRecorder
}

// MockStreamGetterMockRecorder is the mock recorder for MockStreamGetter.
type MockStreamGetterMockRecorder struct {
	mock *MockStreamGetter
}

// NewMockStreamGetter creates a new mock instance.
func NewMockStreamGetter(ctrl *gomock.Controller) *MockStreamGetter {
	mock := &MockStreamGetter{ctrl: ctrl}
	mock.recorder = &MockStreamGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStreamGetter) EXPECT() *MockStreamGetterMockRecorder {
	return m.recorder
}

// GetOrOpenReceiveStream mocks base method.
func (m *MockStreamGetter) GetOrOpenReceiveStream(arg0 protocol.StreamID) (receiveStreamI, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrOpenReceiveStream", arg0)
	ret0, _ := ret[0].(receiveStreamI)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrOpenReceiveStream indicates an expected call of GetOrOpenReceiveStream.
func (mr *MockStreamGetterMockRecorder) GetOrOpenReceiveStream(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrOpenReceiveStream", reflect.TypeOf((*MockStreamGetter)(nil).GetOrOpenReceiveStream), arg0)
}

// GetOrOpenSendStream mocks base method.
func (m *MockStreamGetter) GetOrOpenSendStream(arg0 protocol.StreamID) (sendStreamI, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrOpenSendStream", arg0)
	ret0, _ := ret[0].(sendStreamI)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrOpenSendStream indicates an expected call of GetOrOpenSendStream.
func (mr *MockStreamGetterMockRecorder) GetOrOpenSendStream(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrOpenSendStream", reflect.TypeOf((*MockStreamGetter)(nil).GetOrOpenSendStream), arg0)
}
