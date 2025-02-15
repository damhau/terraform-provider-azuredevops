// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/microsoft/azure-devops-go-api/azuredevops/v6/profile (interfaces: Client)

// Package azdosdkmocks is a generated GoMock package.
package azdosdkmocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	profile "github.com/microsoft/azure-devops-go-api/azuredevops/v6/profile"
)

// MockProfileClient is a mock of Client interface.
type MockProfileClient struct {
	ctrl     *gomock.Controller
	recorder *MockProfileClientMockRecorder
}

// MockProfileClientMockRecorder is the mock recorder for MockProfileClient.
type MockProfileClientMockRecorder struct {
	mock *MockProfileClient
}

// NewMockProfileClient creates a new mock instance.
func NewMockProfileClient(ctrl *gomock.Controller) *MockProfileClient {
	mock := &MockProfileClient{ctrl: ctrl}
	mock.recorder = &MockProfileClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProfileClient) EXPECT() *MockProfileClientMockRecorder {
	return m.recorder
}

// GetProfile mocks base method.
func (m *MockProfileClient) GetProfile(arg0 context.Context, arg1 profile.GetProfileArgs) (*profile.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfile", arg0, arg1)
	ret0, _ := ret[0].(*profile.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfile indicates an expected call of GetProfile.
func (mr *MockProfileClientMockRecorder) GetProfile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfile", reflect.TypeOf((*MockProfileClient)(nil).GetProfile), arg0, arg1)
}
