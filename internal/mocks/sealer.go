// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/pions/quic-go/internal/handshake (interfaces: Sealer)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	protocol "github.com/pions/quic-go/internal/protocol"
)

// MockSealer is a mock of Sealer interface
type MockSealer struct {
	ctrl     *gomock.Controller
	recorder *MockSealerMockRecorder
}

// MockSealerMockRecorder is the mock recorder for MockSealer
type MockSealerMockRecorder struct {
	mock *MockSealer
}

// NewMockSealer creates a new mock instance
func NewMockSealer(ctrl *gomock.Controller) *MockSealer {
	mock := &MockSealer{ctrl: ctrl}
	mock.recorder = &MockSealerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSealer) EXPECT() *MockSealerMockRecorder {
	return m.recorder
}

// EncryptHeader mocks base method
func (m *MockSealer) EncryptHeader(arg0 []byte, arg1 *byte, arg2 []byte) {
	m.ctrl.Call(m, "EncryptHeader", arg0, arg1, arg2)
}

// EncryptHeader indicates an expected call of EncryptHeader
func (mr *MockSealerMockRecorder) EncryptHeader(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EncryptHeader", reflect.TypeOf((*MockSealer)(nil).EncryptHeader), arg0, arg1, arg2)
}

// Overhead mocks base method
func (m *MockSealer) Overhead() int {
	ret := m.ctrl.Call(m, "Overhead")
	ret0, _ := ret[0].(int)
	return ret0
}

// Overhead indicates an expected call of Overhead
func (mr *MockSealerMockRecorder) Overhead() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Overhead", reflect.TypeOf((*MockSealer)(nil).Overhead))
}

// Seal mocks base method
func (m *MockSealer) Seal(arg0, arg1 []byte, arg2 protocol.PacketNumber, arg3 []byte) []byte {
	ret := m.ctrl.Call(m, "Seal", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Seal indicates an expected call of Seal
func (mr *MockSealerMockRecorder) Seal(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Seal", reflect.TypeOf((*MockSealer)(nil).Seal), arg0, arg1, arg2, arg3)
}
