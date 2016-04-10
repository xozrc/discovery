// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/docker/libkv/store (interfaces: Store)

package mock_store

import (
	store "github.com/docker/libkv/store"
	gomock "github.com/golang/mock/gomock"
)

// Mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *_MockStoreRecorder
}

// Recorder for MockStore (not exported)
type _MockStoreRecorder struct {
	mock *MockStore
}

func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &_MockStoreRecorder{mock}
	return mock
}

func (_m *MockStore) EXPECT() *_MockStoreRecorder {
	return _m.recorder
}

func (_m *MockStore) AtomicDelete(_param0 string, _param1 *store.KVPair) (bool, error) {
	ret := _m.ctrl.Call(_m, "AtomicDelete", _param0, _param1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockStoreRecorder) AtomicDelete(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AtomicDelete", arg0, arg1)
}

func (_m *MockStore) AtomicPut(_param0 string, _param1 []byte, _param2 *store.KVPair, _param3 *store.WriteOptions) (bool, *store.KVPair, error) {
	ret := _m.ctrl.Call(_m, "AtomicPut", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*store.KVPair)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockStoreRecorder) AtomicPut(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AtomicPut", arg0, arg1, arg2, arg3)
}

func (_m *MockStore) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockStoreRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockStore) Delete(_param0 string) error {
	ret := _m.ctrl.Call(_m, "Delete", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockStoreRecorder) Delete(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Delete", arg0)
}

func (_m *MockStore) DeleteTree(_param0 string) error {
	ret := _m.ctrl.Call(_m, "DeleteTree", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockStoreRecorder) DeleteTree(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteTree", arg0)
}

func (_m *MockStore) Exists(_param0 string) (bool, error) {
	ret := _m.ctrl.Call(_m, "Exists", _param0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockStoreRecorder) Exists(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Exists", arg0)
}

func (_m *MockStore) Get(_param0 string) (*store.KVPair, error) {
	ret := _m.ctrl.Call(_m, "Get", _param0)
	ret0, _ := ret[0].(*store.KVPair)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockStoreRecorder) Get(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0)
}

func (_m *MockStore) List(_param0 string) ([]*store.KVPair, error) {
	ret := _m.ctrl.Call(_m, "List", _param0)
	ret0, _ := ret[0].([]*store.KVPair)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockStoreRecorder) List(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "List", arg0)
}

func (_m *MockStore) NewLock(_param0 string, _param1 *store.LockOptions) (store.Locker, error) {
	ret := _m.ctrl.Call(_m, "NewLock", _param0, _param1)
	ret0, _ := ret[0].(store.Locker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockStoreRecorder) NewLock(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NewLock", arg0, arg1)
}

func (_m *MockStore) Put(_param0 string, _param1 []byte, _param2 *store.WriteOptions) error {
	ret := _m.ctrl.Call(_m, "Put", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockStoreRecorder) Put(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Put", arg0, arg1, arg2)
}

func (_m *MockStore) Watch(_param0 string, _param1 <-chan struct{}) (<-chan *store.KVPair, error) {
	ret := _m.ctrl.Call(_m, "Watch", _param0, _param1)
	ret0, _ := ret[0].(<-chan *store.KVPair)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockStoreRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Watch", arg0, arg1)
}

func (_m *MockStore) WatchTree(_param0 string, _param1 <-chan struct{}) (<-chan []*store.KVPair, error) {
	ret := _m.ctrl.Call(_m, "WatchTree", _param0, _param1)
	ret0, _ := ret[0].(<-chan []*store.KVPair)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockStoreRecorder) WatchTree(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WatchTree", arg0, arg1)
}
