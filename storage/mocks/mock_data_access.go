// SPDX-License-Identifier: ISC
// Code generated by MockGen. DO NOT EDIT.
// Source: data_access.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	iterator "github.com/syndtr/goleveldb/leveldb/iterator"
	util "github.com/syndtr/goleveldb/leveldb/util"
)

// MockDataAccess is a mock of DataAccess interface
type MockDataAccess struct {
	ctrl     *gomock.Controller
	recorder *MockDataAccessMockRecorder
}

// MockDataAccessMockRecorder is the mock recorder for MockDataAccess
type MockDataAccessMockRecorder struct {
	mock *MockDataAccess
}

// NewMockDataAccess creates a new mock instance
func NewMockDataAccess(ctrl *gomock.Controller) *MockDataAccess {
	mock := &MockDataAccess{ctrl: ctrl}
	mock.recorder = &MockDataAccessMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDataAccess) EXPECT() *MockDataAccessMockRecorder {
	return m.recorder
}

// Abort mocks base method
func (m *MockDataAccess) Abort() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Abort")
}

// Abort indicates an expected call of Abort
func (mr *MockDataAccessMockRecorder) Abort() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Abort", reflect.TypeOf((*MockDataAccess)(nil).Abort))
}

// Begin mocks base method
func (m *MockDataAccess) Begin() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Begin")
	ret0, _ := ret[0].(error)
	return ret0
}

// Begin indicates an expected call of Begin
func (mr *MockDataAccessMockRecorder) Begin() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockDataAccess)(nil).Begin))
}

// Commit mocks base method
func (m *MockDataAccess) Commit() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit
func (mr *MockDataAccessMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockDataAccess)(nil).Commit))
}

// Delete mocks base method
func (m *MockDataAccess) Delete(arg0 []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", arg0)
}

// Delete indicates an expected call of Delete
func (mr *MockDataAccessMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDataAccess)(nil).Delete), arg0)
}

// DumpTx mocks base method
func (m *MockDataAccess) DumpTx() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DumpTx")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// DumpTx indicates an expected call of DumpTx
func (mr *MockDataAccessMockRecorder) DumpTx() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DumpTx", reflect.TypeOf((*MockDataAccess)(nil).DumpTx))
}

// Get mocks base method
func (m *MockDataAccess) Get(arg0 []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockDataAccessMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDataAccess)(nil).Get), arg0)
}

// Has mocks base method
func (m *MockDataAccess) Has(arg0 []byte) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Has indicates an expected call of Has
func (mr *MockDataAccessMockRecorder) Has(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockDataAccess)(nil).Has), arg0)
}

// InUse mocks base method
func (m *MockDataAccess) InUse() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InUse")
	ret0, _ := ret[0].(bool)
	return ret0
}

// InUse indicates an expected call of InUse
func (mr *MockDataAccessMockRecorder) InUse() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InUse", reflect.TypeOf((*MockDataAccess)(nil).InUse))
}

// Iterator mocks base method
func (m *MockDataAccess) Iterator(arg0 *util.Range) iterator.Iterator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Iterator", arg0)
	ret0, _ := ret[0].(iterator.Iterator)
	return ret0
}

// Iterator indicates an expected call of Iterator
func (mr *MockDataAccessMockRecorder) Iterator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iterator", reflect.TypeOf((*MockDataAccess)(nil).Iterator), arg0)
}

// Put mocks base method
func (m *MockDataAccess) Put(arg0, arg1 []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Put", arg0, arg1)
}

// Put indicates an expected call of Put
func (mr *MockDataAccessMockRecorder) Put(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockDataAccess)(nil).Put), arg0, arg1)
}
