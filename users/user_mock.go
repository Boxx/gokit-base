// Automatically generated by MockGen. DO NOT EDIT!
// Source: users/user.go

package users

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *_MockRepositoryRecorder
}

// Recorder for MockRepository (not exported)
type _MockRepositoryRecorder struct {
	mock *MockRepository
}

func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &_MockRepositoryRecorder{mock}
	return mock
}

func (_m *MockRepository) EXPECT() *_MockRepositoryRecorder {
	return _m.recorder
}

func (_m *MockRepository) Store(user *User) error {
	ret := _m.ctrl.Call(_m, "Store", user)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockRepositoryRecorder) Store(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Store", arg0)
}

func (_m *MockRepository) Find(id int) (*User, error) {
	ret := _m.ctrl.Call(_m, "Find", id)
	ret0, _ := ret[0].(*User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockRepositoryRecorder) Find(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Find", arg0)
}

func (_m *MockRepository) FindAll() []*User {
	ret := _m.ctrl.Call(_m, "FindAll")
	ret0, _ := ret[0].([]*User)
	return ret0
}

func (_mr *_MockRepositoryRecorder) FindAll() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FindAll")
}
