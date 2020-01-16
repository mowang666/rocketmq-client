/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by MockGen. DO NOT EDIT.
// Source: remote_client.go

// Package remote is a generated GoMock package.
package remote

import (
	context "context"
	primitive "github.com/mowang666/rocketmq-client/primitive"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRemotingClient is a mock of RemotingClient interface
type MockRemotingClient struct {
	ctrl     *gomock.Controller
	recorder *MockRemotingClientMockRecorder
}

// MockRemotingClientMockRecorder is the mock recorder for MockRemotingClient
type MockRemotingClientMockRecorder struct {
	mock *MockRemotingClient
}

// NewMockRemotingClient creates a new mock instance
func NewMockRemotingClient(ctrl *gomock.Controller) *MockRemotingClient {
	mock := &MockRemotingClient{ctrl: ctrl}
	mock.recorder = &MockRemotingClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRemotingClient) EXPECT() *MockRemotingClientMockRecorder {
	return m.recorder
}

// RegisterRequestFunc mocks base method
func (m *MockRemotingClient) RegisterRequestFunc(code int16, f ClientRequestFunc) {
	m.ctrl.Call(m, "RegisterRequestFunc", code, f)
}

// RegisterRequestFunc indicates an expected call of RegisterRequestFunc
func (mr *MockRemotingClientMockRecorder) RegisterRequestFunc(code, f interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterRequestFunc", reflect.TypeOf((*MockRemotingClient)(nil).RegisterRequestFunc), code, f)
}

// RegisterInterceptor mocks base method
func (m *MockRemotingClient) RegisterInterceptor(interceptors ...primitive.Interceptor) {
	varargs := []interface{}{}
	for _, a := range interceptors {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "RegisterInterceptor", varargs...)
}

// RegisterInterceptor indicates an expected call of RegisterInterceptor
func (mr *MockRemotingClientMockRecorder) RegisterInterceptor(interceptors ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterInterceptor", reflect.TypeOf((*MockRemotingClient)(nil).RegisterInterceptor), interceptors...)
}

// InvokeSync mocks base method
func (m *MockRemotingClient) InvokeSync(ctx context.Context, addr string, request *RemotingCommand) (*RemotingCommand, error) {
	ret := m.ctrl.Call(m, "InvokeSync", ctx, addr, request)
	ret0, _ := ret[0].(*RemotingCommand)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InvokeSync indicates an expected call of InvokeSync
func (mr *MockRemotingClientMockRecorder) InvokeSync(ctx, addr, request interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvokeSync", reflect.TypeOf((*MockRemotingClient)(nil).InvokeSync), ctx, addr, request)
}

// InvokeAsync mocks base method
func (m *MockRemotingClient) InvokeAsync(ctx context.Context, addr string, request *RemotingCommand, callback func(*ResponseFuture)) error {
	ret := m.ctrl.Call(m, "InvokeAsync", ctx, addr, request, callback)
	ret0, _ := ret[0].(error)
	return ret0
}

// InvokeAsync indicates an expected call of InvokeAsync
func (mr *MockRemotingClientMockRecorder) InvokeAsync(ctx, addr, request, callback interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvokeAsync", reflect.TypeOf((*MockRemotingClient)(nil).InvokeAsync), ctx, addr, request, callback)
}

// InvokeOneWay mocks base method
func (m *MockRemotingClient) InvokeOneWay(ctx context.Context, addr string, request *RemotingCommand) error {
	ret := m.ctrl.Call(m, "InvokeOneWay", ctx, addr, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// InvokeOneWay indicates an expected call of InvokeOneWay
func (mr *MockRemotingClientMockRecorder) InvokeOneWay(ctx, addr, request interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvokeOneWay", reflect.TypeOf((*MockRemotingClient)(nil).InvokeOneWay), ctx, addr, request)
}

// ShutDown mocks base method
func (m *MockRemotingClient) ShutDown() {
	m.ctrl.Call(m, "ShutDown")
}

// ShutDown indicates an expected call of ShutDown
func (mr *MockRemotingClientMockRecorder) ShutDown() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShutDown", reflect.TypeOf((*MockRemotingClient)(nil).ShutDown))
}
