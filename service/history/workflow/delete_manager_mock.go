// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by MockGen. DO NOT EDIT.
// Source: delete_manager.go

// Package workflow is a generated GoMock package.
package workflow

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "go.temporal.io/api/common/v1"
	namespace "go.temporal.io/server/common/namespace"
)

// MockDeleteManager is a mock of DeleteManager interface.
type MockDeleteManager struct {
	ctrl     *gomock.Controller
	recorder *MockDeleteManagerMockRecorder
}

// MockDeleteManagerMockRecorder is the mock recorder for MockDeleteManager.
type MockDeleteManagerMockRecorder struct {
	mock *MockDeleteManager
}

// NewMockDeleteManager creates a new mock instance.
func NewMockDeleteManager(ctrl *gomock.Controller) *MockDeleteManager {
	mock := &MockDeleteManager{ctrl: ctrl}
	mock.recorder = &MockDeleteManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeleteManager) EXPECT() *MockDeleteManagerMockRecorder {
	return m.recorder
}

// AddDeleteWorkflowExecutionTask mocks base method.
func (m *MockDeleteManager) AddDeleteWorkflowExecutionTask(ctx context.Context, nsID namespace.ID, we v1.WorkflowExecution, ms MutableState, workflowClosedVersion int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDeleteWorkflowExecutionTask", ctx, nsID, we, ms, workflowClosedVersion)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddDeleteWorkflowExecutionTask indicates an expected call of AddDeleteWorkflowExecutionTask.
func (mr *MockDeleteManagerMockRecorder) AddDeleteWorkflowExecutionTask(ctx, nsID, we, ms, workflowClosedVersion interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDeleteWorkflowExecutionTask", reflect.TypeOf((*MockDeleteManager)(nil).AddDeleteWorkflowExecutionTask), ctx, nsID, we, ms, workflowClosedVersion)
}

// DeleteWorkflowExecution mocks base method.
func (m *MockDeleteManager) DeleteWorkflowExecution(ctx context.Context, nsID namespace.ID, we v1.WorkflowExecution, weCtx Context, ms MutableState, forceDeleteFromOpenVisibility bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWorkflowExecution", ctx, nsID, we, weCtx, ms, forceDeleteFromOpenVisibility)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWorkflowExecution indicates an expected call of DeleteWorkflowExecution.
func (mr *MockDeleteManagerMockRecorder) DeleteWorkflowExecution(ctx, nsID, we, weCtx, ms, forceDeleteFromOpenVisibility interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWorkflowExecution", reflect.TypeOf((*MockDeleteManager)(nil).DeleteWorkflowExecution), ctx, nsID, we, weCtx, ms, forceDeleteFromOpenVisibility)
}

// DeleteWorkflowExecutionByRetention mocks base method.
func (m *MockDeleteManager) DeleteWorkflowExecutionByRetention(ctx context.Context, nsID namespace.ID, we v1.WorkflowExecution, weCtx Context, ms MutableState, archiveIfEnabled bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWorkflowExecutionByRetention", ctx, nsID, we, weCtx, ms, archiveIfEnabled)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWorkflowExecutionByRetention indicates an expected call of DeleteWorkflowExecutionByRetention.
func (mr *MockDeleteManagerMockRecorder) DeleteWorkflowExecutionByRetention(ctx, nsID, we, weCtx, ms, archiveIfEnabled interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWorkflowExecutionByRetention", reflect.TypeOf((*MockDeleteManager)(nil).DeleteWorkflowExecutionByRetention), ctx, nsID, we, weCtx, ms, archiveIfEnabled)
}
