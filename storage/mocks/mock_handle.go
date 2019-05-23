// Code generated by MockGen. DO NOT EDIT.
// Source: handle.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockHandle is a mock of Handle interface
type MockHandle struct {
	ctrl     *gomock.Controller
	recorder *MockHandleMockRecorder
}

// MockHandleMockRecorder is the mock recorder for MockHandle
type MockHandleMockRecorder struct {
	mock *MockHandle
}

// NewMockHandle creates a new mock instance
func NewMockHandle(ctrl *gomock.Controller) *MockHandle {
	mock := &MockHandle{ctrl: ctrl}
	mock.recorder = &MockHandleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHandle) EXPECT() *MockHandleMockRecorder {
	return m.recorder
}

// Put mocks base method
func (m *MockHandle) Put(key, value []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Put", key, value)
}

// Put indicates an expected call of Put
func (mr *MockHandleMockRecorder) Put(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockHandle)(nil).Put), key, value)
}

// put mocks base method
func (m *MockHandle) put(key, value []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "put", key, value)
}

// put indicates an expected call of put
func (mr *MockHandleMockRecorder) put(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "put", reflect.TypeOf((*MockHandle)(nil).put), key, value)
}

// PutN mocks base method
func (m *MockHandle) PutN(key []byte, value uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PutN", key, value)
}

// PutN indicates an expected call of PutN
func (mr *MockHandleMockRecorder) PutN(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutN", reflect.TypeOf((*MockHandle)(nil).PutN), key, value)
}

// putN mocks base method
func (m *MockHandle) putN(key []byte, value uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "putN", key, value)
}

// putN indicates an expected call of putN
func (mr *MockHandleMockRecorder) putN(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "putN", reflect.TypeOf((*MockHandle)(nil).putN), key, value)
}

// Delete mocks base method
func (m *MockHandle) Delete(key []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", key)
}

// Delete indicates an expected call of Delete
func (mr *MockHandleMockRecorder) Delete(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockHandle)(nil).Delete), key)
}

// remove mocks base method
func (m *MockHandle) remove(key []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "remove", key)
}

// remove indicates an expected call of remove
func (mr *MockHandleMockRecorder) remove(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "remove", reflect.TypeOf((*MockHandle)(nil).remove), key)
}

// Get mocks base method
func (m *MockHandle) Get(key []byte) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockHandleMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockHandle)(nil).Get), key)
}

// GetN mocks base method
func (m *MockHandle) GetN(key []byte) (uint64, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetN", key)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetN indicates an expected call of GetN
func (mr *MockHandleMockRecorder) GetN(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetN", reflect.TypeOf((*MockHandle)(nil).GetN), key)
}

// getN mocks base method
func (m *MockHandle) getN(key []byte) (uint64, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getN", key)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// getN indicates an expected call of getN
func (mr *MockHandleMockRecorder) getN(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getN", reflect.TypeOf((*MockHandle)(nil).getN), key)
}

// GetNB mocks base method
func (m *MockHandle) GetNB(key []byte) (uint64, []byte) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNB", key)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].([]byte)
	return ret0, ret1
}

// GetNB indicates an expected call of GetNB
func (mr *MockHandleMockRecorder) GetNB(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNB", reflect.TypeOf((*MockHandle)(nil).GetNB), key)
}

// getNB mocks base method
func (m *MockHandle) getNB(key []byte) (uint64, []byte) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getNB", key)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].([]byte)
	return ret0, ret1
}

// getNB indicates an expected call of getNB
func (mr *MockHandleMockRecorder) getNB(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getNB", reflect.TypeOf((*MockHandle)(nil).getNB), key)
}

// Has mocks base method
func (m *MockHandle) Has(key []byte) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", key)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has
func (mr *MockHandleMockRecorder) Has(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockHandle)(nil).Has), key)
}

// Begin mocks base method
func (m *MockHandle) Begin() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Begin")
}

// Begin indicates an expected call of Begin
func (mr *MockHandleMockRecorder) Begin() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockHandle)(nil).Begin))
}

// Commit mocks base method
func (m *MockHandle) Commit() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit
func (mr *MockHandleMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockHandle)(nil).Commit))
}
