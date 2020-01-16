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
// Source: offset_store.go

// Package consumer is a generated GoMock package.
package consumer

import (
	primitive "github.com/mowang666/rocketmq-client/primitive"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockOffsetStore is a mock of OffsetStore interface
type MockOffsetStore struct {
	ctrl     *gomock.Controller
	recorder *MockOffsetStoreMockRecorder
}

// MockOffsetStoreMockRecorder is the mock recorder for MockOffsetStore
type MockOffsetStoreMockRecorder struct {
	mock *MockOffsetStore
}

// NewMockOffsetStore creates a new mock instance
func NewMockOffsetStore(ctrl *gomock.Controller) *MockOffsetStore {
	mock := &MockOffsetStore{ctrl: ctrl}
	mock.recorder = &MockOffsetStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOffsetStore) EXPECT() *MockOffsetStoreMockRecorder {
	return m.recorder
}

// persist mocks base method
func (m *MockOffsetStore) persist(mqs []*primitive.MessageQueue) {
	m.ctrl.Call(m, "persist", mqs)
}

// persist indicates an expected call of persist
func (mr *MockOffsetStoreMockRecorder) persist(mqs interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "persist", reflect.TypeOf((*MockOffsetStore)(nil).persist), mqs)
}

// remove mocks base method
func (m *MockOffsetStore) remove(mq *primitive.MessageQueue) {
	m.ctrl.Call(m, "remove", mq)
}

// remove indicates an expected call of remove
func (mr *MockOffsetStoreMockRecorder) remove(mq interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "remove", reflect.TypeOf((*MockOffsetStore)(nil).remove), mq)
}

// read mocks base method
func (m *MockOffsetStore) read(mq *primitive.MessageQueue, t readType) int64 {
	ret := m.ctrl.Call(m, "read", mq, t)
	ret0, _ := ret[0].(int64)
	return ret0
}

// read indicates an expected call of read
func (mr *MockOffsetStoreMockRecorder) read(mq, t interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "read", reflect.TypeOf((*MockOffsetStore)(nil).read), mq, t)
}

// update mocks base method
func (m *MockOffsetStore) update(mq *primitive.MessageQueue, offset int64, increaseOnly bool) {
	m.ctrl.Call(m, "update", mq, offset, increaseOnly)
}

// update indicates an expected call of update
func (mr *MockOffsetStoreMockRecorder) update(mq, offset, increaseOnly interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "update", reflect.TypeOf((*MockOffsetStore)(nil).update), mq, offset, increaseOnly)
}
