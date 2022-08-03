// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/choria-io/go-choria/aagent/watchers (interfaces: Machine)

//  mockgen -package watchers -destination machine_mock_test.go github.com/choria-io/go-choria/aagent/watchers Machine

// Package watchers is a generated GoMock package.
package watchers

import (
	json "encoding/json"
	reflect "reflect"

	lifecycle "github.com/choria-io/go-choria/lifecycle"
	gomock "github.com/golang/mock/gomock"
	jsm_go "github.com/nats-io/jsm.go"
)

// MockMachine is a mock of Machine interface.
type MockMachine struct {
	ctrl     *gomock.Controller
	recorder *MockMachineMockRecorder
}

// MockMachineMockRecorder is the mock recorder for MockMachine.
type MockMachineMockRecorder struct {
	mock *MockMachine
}

// NewMockMachine creates a new mock instance.
func NewMockMachine(ctrl *gomock.Controller) *MockMachine {
	mock := &MockMachine{ctrl: ctrl}
	mock.recorder = &MockMachineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMachine) EXPECT() *MockMachineMockRecorder {
	return m.recorder
}

// ChoriaStatusFile mocks base method.
func (m *MockMachine) ChoriaStatusFile() (string, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChoriaStatusFile")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int)
	return ret0, ret1
}

// ChoriaStatusFile indicates an expected call of ChoriaStatusFile.
func (mr *MockMachineMockRecorder) ChoriaStatusFile() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChoriaStatusFile", reflect.TypeOf((*MockMachine)(nil).ChoriaStatusFile))
}

// Data mocks base method.
func (m *MockMachine) Data() map[string]any {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Data")
	ret0, _ := ret[0].(map[string]any)
	return ret0
}

// Data indicates an expected call of Data.
func (mr *MockMachineMockRecorder) Data() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Data", reflect.TypeOf((*MockMachine)(nil).Data))
}

// DataDelete mocks base method.
func (m *MockMachine) DataDelete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataDelete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DataDelete indicates an expected call of DataDelete.
func (mr *MockMachineMockRecorder) DataDelete(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataDelete", reflect.TypeOf((*MockMachine)(nil).DataDelete), arg0)
}

// DataGet mocks base method.
func (m *MockMachine) DataGet(arg0 string) (any, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataGet", arg0)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// DataGet indicates an expected call of DataGet.
func (mr *MockMachineMockRecorder) DataGet(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataGet", reflect.TypeOf((*MockMachine)(nil).DataGet), arg0)
}

// DataPut mocks base method.
func (m *MockMachine) DataPut(arg0 string, arg1 any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataPut", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DataPut indicates an expected call of DataPut.
func (mr *MockMachineMockRecorder) DataPut(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataPut", reflect.TypeOf((*MockMachine)(nil).DataPut), arg0, arg1)
}

// Debugf mocks base method.
func (m *MockMachine) Debugf(arg0, arg1 string, arg2 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf.
func (mr *MockMachineMockRecorder) Debugf(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*MockMachine)(nil).Debugf), varargs...)
}

// Directory mocks base method.
func (m *MockMachine) Directory() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Directory")
	ret0, _ := ret[0].(string)
	return ret0
}

// Directory indicates an expected call of Directory.
func (mr *MockMachineMockRecorder) Directory() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Directory", reflect.TypeOf((*MockMachine)(nil).Directory))
}

// Errorf mocks base method.
func (m *MockMachine) Errorf(arg0, arg1 string, arg2 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf.
func (mr *MockMachineMockRecorder) Errorf(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockMachine)(nil).Errorf), varargs...)
}

// Facts mocks base method.
func (m *MockMachine) Facts() json.RawMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Facts")
	ret0, _ := ret[0].(json.RawMessage)
	return ret0
}

// Facts indicates an expected call of Facts.
func (mr *MockMachineMockRecorder) Facts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Facts", reflect.TypeOf((*MockMachine)(nil).Facts))
}

// Identity mocks base method.
func (m *MockMachine) Identity() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Identity")
	ret0, _ := ret[0].(string)
	return ret0
}

// Identity indicates an expected call of Identity.
func (mr *MockMachineMockRecorder) Identity() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Identity", reflect.TypeOf((*MockMachine)(nil).Identity))
}

// Infof mocks base method.
func (m *MockMachine) Infof(arg0, arg1 string, arg2 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infof", varargs...)
}

// Infof indicates an expected call of Infof.
func (mr *MockMachineMockRecorder) Infof(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infof", reflect.TypeOf((*MockMachine)(nil).Infof), varargs...)
}

// InstanceID mocks base method.
func (m *MockMachine) InstanceID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstanceID")
	ret0, _ := ret[0].(string)
	return ret0
}

// InstanceID indicates an expected call of InstanceID.
func (mr *MockMachineMockRecorder) InstanceID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceID", reflect.TypeOf((*MockMachine)(nil).InstanceID))
}

// JetStreamConnection mocks base method.
func (m *MockMachine) JetStreamConnection() (*jsm_go.Manager, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JetStreamConnection")
	ret0, _ := ret[0].(*jsm_go.Manager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// JetStreamConnection indicates an expected call of JetStreamConnection.
func (mr *MockMachineMockRecorder) JetStreamConnection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JetStreamConnection", reflect.TypeOf((*MockMachine)(nil).JetStreamConnection))
}

// MainCollective mocks base method.
func (m *MockMachine) MainCollective() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MainCollective")
	ret0, _ := ret[0].(string)
	return ret0
}

// MainCollective indicates an expected call of MainCollective.
func (mr *MockMachineMockRecorder) MainCollective() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MainCollective", reflect.TypeOf((*MockMachine)(nil).MainCollective))
}

// Name mocks base method.
func (m *MockMachine) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockMachineMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockMachine)(nil).Name))
}

// NotifyWatcherState mocks base method.
func (m *MockMachine) NotifyWatcherState(arg0 string, arg1 any) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NotifyWatcherState", arg0, arg1)
}

// NotifyWatcherState indicates an expected call of NotifyWatcherState.
func (mr *MockMachineMockRecorder) NotifyWatcherState(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyWatcherState", reflect.TypeOf((*MockMachine)(nil).NotifyWatcherState), arg0, arg1)
}

// OverrideData mocks base method.
func (m *MockMachine) OverrideData() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OverrideData")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OverrideData indicates an expected call of OverrideData.
func (mr *MockMachineMockRecorder) OverrideData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OverrideData", reflect.TypeOf((*MockMachine)(nil).OverrideData))
}

// PublishLifecycleEvent mocks base method.
func (m *MockMachine) PublishLifecycleEvent(arg0 lifecycle.Type, arg1 ...lifecycle.Option) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "PublishLifecycleEvent", varargs...)
}

// PublishLifecycleEvent indicates an expected call of PublishLifecycleEvent.
func (mr *MockMachineMockRecorder) PublishLifecycleEvent(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishLifecycleEvent", reflect.TypeOf((*MockMachine)(nil).PublishLifecycleEvent), varargs...)
}

// State mocks base method.
func (m *MockMachine) State() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "State")
	ret0, _ := ret[0].(string)
	return ret0
}

// State indicates an expected call of State.
func (mr *MockMachineMockRecorder) State() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "State", reflect.TypeOf((*MockMachine)(nil).State))
}

// TextFileDirectory mocks base method.
func (m *MockMachine) TextFileDirectory() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TextFileDirectory")
	ret0, _ := ret[0].(string)
	return ret0
}

// TextFileDirectory indicates an expected call of TextFileDirectory.
func (mr *MockMachineMockRecorder) TextFileDirectory() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TextFileDirectory", reflect.TypeOf((*MockMachine)(nil).TextFileDirectory))
}

// TimeStampSeconds mocks base method.
func (m *MockMachine) TimeStampSeconds() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TimeStampSeconds")
	ret0, _ := ret[0].(int64)
	return ret0
}

// TimeStampSeconds indicates an expected call of TimeStampSeconds.
func (mr *MockMachineMockRecorder) TimeStampSeconds() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TimeStampSeconds", reflect.TypeOf((*MockMachine)(nil).TimeStampSeconds))
}

// Transition mocks base method.
func (m *MockMachine) Transition(arg0 string, arg1 ...any) error {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Transition", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Transition indicates an expected call of Transition.
func (mr *MockMachineMockRecorder) Transition(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transition", reflect.TypeOf((*MockMachine)(nil).Transition), varargs...)
}

// Version mocks base method.
func (m *MockMachine) Version() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(string)
	return ret0
}

// Version indicates an expected call of Version.
func (mr *MockMachineMockRecorder) Version() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockMachine)(nil).Version))
}

// Warnf mocks base method.
func (m *MockMachine) Warnf(arg0, arg1 string, arg2 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warnf", varargs...)
}

// Warnf indicates an expected call of Warnf.
func (mr *MockMachineMockRecorder) Warnf(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warnf", reflect.TypeOf((*MockMachine)(nil).Warnf), varargs...)
}

// Watchers mocks base method.
func (m *MockMachine) Watchers() []*WatcherDef {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watchers")
	ret0, _ := ret[0].([]*WatcherDef)
	return ret0
}

// Watchers indicates an expected call of Watchers.
func (mr *MockMachineMockRecorder) Watchers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watchers", reflect.TypeOf((*MockMachine)(nil).Watchers))
}
