// Copyright 2023 API7.ai, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: ./types.go

// Package cloud is a generated GoMock package.
package cloud

import (
	reflect "reflect"

	cloud_go_sdk "github.com/api7/cloud-go-sdk"
	gomock "github.com/golang/mock/gomock"
)

// MockAPI is a mock of API interface.
type MockAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAPIMockRecorder
}

// MockAPIMockRecorder is the mock recorder for MockAPI.
type MockAPIMockRecorder struct {
	mock *MockAPI
}

// NewMockAPI creates a new mock instance.
func NewMockAPI(ctrl *gomock.Controller) *MockAPI {
	mock := &MockAPI{ctrl: ctrl}
	mock.recorder = &MockAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPI) EXPECT() *MockAPIMockRecorder {
	return m.recorder
}

// CreateConsumer mocks base method.
func (m *MockAPI) CreateConsumer(clusterID cloud_go_sdk.ID, consumer *cloud_go_sdk.Consumer) (*cloud_go_sdk.Consumer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateConsumer", clusterID, consumer)
	ret0, _ := ret[0].(*cloud_go_sdk.Consumer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateConsumer indicates an expected call of CreateConsumer.
func (mr *MockAPIMockRecorder) CreateConsumer(clusterID, consumer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateConsumer", reflect.TypeOf((*MockAPI)(nil).CreateConsumer), clusterID, consumer)
}

// CreateSSL mocks base method.
func (m *MockAPI) CreateSSL(clusterID cloud_go_sdk.ID, ssl *cloud_go_sdk.Certificate) (*cloud_go_sdk.CertificateDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSSL", clusterID, ssl)
	ret0, _ := ret[0].(*cloud_go_sdk.CertificateDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSSL indicates an expected call of CreateSSL.
func (mr *MockAPIMockRecorder) CreateSSL(clusterID, ssl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSSL", reflect.TypeOf((*MockAPI)(nil).CreateSSL), clusterID, ssl)
}

// CreateService mocks base method.
func (m *MockAPI) CreateService(clusterID cloud_go_sdk.ID, svc *cloud_go_sdk.Application) (*cloud_go_sdk.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateService", clusterID, svc)
	ret0, _ := ret[0].(*cloud_go_sdk.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateService indicates an expected call of CreateService.
func (mr *MockAPIMockRecorder) CreateService(clusterID, svc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateService", reflect.TypeOf((*MockAPI)(nil).CreateService), clusterID, svc)
}

// DebugShowConfig mocks base method.
func (m *MockAPI) DebugShowConfig(clusterID cloud_go_sdk.ID, resource string, id cloud_go_sdk.ID) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DebugShowConfig", clusterID, resource, id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DebugShowConfig indicates an expected call of DebugShowConfig.
func (mr *MockAPIMockRecorder) DebugShowConfig(clusterID, resource, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DebugShowConfig", reflect.TypeOf((*MockAPI)(nil).DebugShowConfig), clusterID, resource, id)
}

// DeleteAPI mocks base method.
func (m *MockAPI) DeleteAPI(clusterID, appID, apiID cloud_go_sdk.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAPI", clusterID, appID, apiID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAPI indicates an expected call of DeleteAPI.
func (mr *MockAPIMockRecorder) DeleteAPI(clusterID, appID, apiID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAPI", reflect.TypeOf((*MockAPI)(nil).DeleteAPI), clusterID, appID, apiID)
}

// DeleteConsumer mocks base method.
func (m *MockAPI) DeleteConsumer(clusterID, consumerID cloud_go_sdk.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteConsumer", clusterID, consumerID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteConsumer indicates an expected call of DeleteConsumer.
func (mr *MockAPIMockRecorder) DeleteConsumer(clusterID, consumerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteConsumer", reflect.TypeOf((*MockAPI)(nil).DeleteConsumer), clusterID, consumerID)
}

// DeleteSSL mocks base method.
func (m *MockAPI) DeleteSSL(clusterID, sslID cloud_go_sdk.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSSL", clusterID, sslID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSSL indicates an expected call of DeleteSSL.
func (mr *MockAPIMockRecorder) DeleteSSL(clusterID, sslID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSSL", reflect.TypeOf((*MockAPI)(nil).DeleteSSL), clusterID, sslID)
}

// DeleteService mocks base method.
func (m *MockAPI) DeleteService(clusterID, appID cloud_go_sdk.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteService", clusterID, appID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteService indicates an expected call of DeleteService.
func (mr *MockAPIMockRecorder) DeleteService(clusterID, appID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteService", reflect.TypeOf((*MockAPI)(nil).DeleteService), clusterID, appID)
}

// GetCloudLuaModule mocks base method.
func (m *MockAPI) GetCloudLuaModule() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCloudLuaModule")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCloudLuaModule indicates an expected call of GetCloudLuaModule.
func (mr *MockAPIMockRecorder) GetCloudLuaModule() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCloudLuaModule", reflect.TypeOf((*MockAPI)(nil).GetCloudLuaModule))
}

// GetClusterDetail mocks base method.
func (m *MockAPI) GetClusterDetail(clusterID cloud_go_sdk.ID) (*cloud_go_sdk.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterDetail", clusterID)
	ret0, _ := ret[0].(*cloud_go_sdk.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterDetail indicates an expected call of GetClusterDetail.
func (mr *MockAPIMockRecorder) GetClusterDetail(clusterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterDetail", reflect.TypeOf((*MockAPI)(nil).GetClusterDetail), clusterID)
}

// GetConsumer mocks base method.
func (m *MockAPI) GetConsumer(clusterID, consumerID cloud_go_sdk.ID) (*cloud_go_sdk.Consumer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConsumer", clusterID, consumerID)
	ret0, _ := ret[0].(*cloud_go_sdk.Consumer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConsumer indicates an expected call of GetConsumer.
func (mr *MockAPIMockRecorder) GetConsumer(clusterID, consumerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConsumer", reflect.TypeOf((*MockAPI)(nil).GetConsumer), clusterID, consumerID)
}

// GetDefaultCluster mocks base method.
func (m *MockAPI) GetDefaultCluster() (*cloud_go_sdk.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDefaultCluster")
	ret0, _ := ret[0].(*cloud_go_sdk.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDefaultCluster indicates an expected call of GetDefaultCluster.
func (mr *MockAPIMockRecorder) GetDefaultCluster() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDefaultCluster", reflect.TypeOf((*MockAPI)(nil).GetDefaultCluster))
}

// GetDefaultOrganization mocks base method.
func (m *MockAPI) GetDefaultOrganization() (*cloud_go_sdk.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDefaultOrganization")
	ret0, _ := ret[0].(*cloud_go_sdk.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDefaultOrganization indicates an expected call of GetDefaultOrganization.
func (mr *MockAPIMockRecorder) GetDefaultOrganization() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDefaultOrganization", reflect.TypeOf((*MockAPI)(nil).GetDefaultOrganization))
}

// GetSSL mocks base method.
func (m *MockAPI) GetSSL(clusterID, sslID cloud_go_sdk.ID) (*cloud_go_sdk.CertificateDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSSL", clusterID, sslID)
	ret0, _ := ret[0].(*cloud_go_sdk.CertificateDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSSL indicates an expected call of GetSSL.
func (mr *MockAPIMockRecorder) GetSSL(clusterID, sslID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSSL", reflect.TypeOf((*MockAPI)(nil).GetSSL), clusterID, sslID)
}

// GetService mocks base method.
func (m *MockAPI) GetService(clusterID, appID cloud_go_sdk.ID) (*cloud_go_sdk.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetService", clusterID, appID)
	ret0, _ := ret[0].(*cloud_go_sdk.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetService indicates an expected call of GetService.
func (mr *MockAPIMockRecorder) GetService(clusterID, appID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetService", reflect.TypeOf((*MockAPI)(nil).GetService), clusterID, appID)
}

// GetStartupConfig mocks base method.
func (m *MockAPI) GetStartupConfig(clusterID cloud_go_sdk.ID, configType StartupConfigType) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStartupConfig", clusterID, configType)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStartupConfig indicates an expected call of GetStartupConfig.
func (mr *MockAPIMockRecorder) GetStartupConfig(clusterID, configType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStartupConfig", reflect.TypeOf((*MockAPI)(nil).GetStartupConfig), clusterID, configType)
}

// GetTLSBundle mocks base method.
func (m *MockAPI) GetTLSBundle(clusterID cloud_go_sdk.ID) (*cloud_go_sdk.TLSBundle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTLSBundle", clusterID)
	ret0, _ := ret[0].(*cloud_go_sdk.TLSBundle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTLSBundle indicates an expected call of GetTLSBundle.
func (mr *MockAPIMockRecorder) GetTLSBundle(clusterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTLSBundle", reflect.TypeOf((*MockAPI)(nil).GetTLSBundle), clusterID)
}

// ListClusters mocks base method.
func (m *MockAPI) ListClusters(orgID cloud_go_sdk.ID, limit, skip int) ([]*cloud_go_sdk.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListClusters", orgID, limit, skip)
	ret0, _ := ret[0].([]*cloud_go_sdk.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListClusters indicates an expected call of ListClusters.
func (mr *MockAPIMockRecorder) ListClusters(orgID, limit, skip interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClusters", reflect.TypeOf((*MockAPI)(nil).ListClusters), orgID, limit, skip)
}

// ListConsumers mocks base method.
func (m *MockAPI) ListConsumers(clusterID cloud_go_sdk.ID, limit, skip int) ([]*cloud_go_sdk.Consumer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListConsumers", clusterID, limit, skip)
	ret0, _ := ret[0].([]*cloud_go_sdk.Consumer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListConsumers indicates an expected call of ListConsumers.
func (mr *MockAPIMockRecorder) ListConsumers(clusterID, limit, skip interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListConsumers", reflect.TypeOf((*MockAPI)(nil).ListConsumers), clusterID, limit, skip)
}

// ListSSL mocks base method.
func (m *MockAPI) ListSSL(clusterID cloud_go_sdk.ID, limit, skip int) ([]*cloud_go_sdk.CertificateDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSSL", clusterID, limit, skip)
	ret0, _ := ret[0].([]*cloud_go_sdk.CertificateDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSSL indicates an expected call of ListSSL.
func (mr *MockAPIMockRecorder) ListSSL(clusterID, limit, skip interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSSL", reflect.TypeOf((*MockAPI)(nil).ListSSL), clusterID, limit, skip)
}

// ListServices mocks base method.
func (m *MockAPI) ListServices(clusterID cloud_go_sdk.ID, limit, skip int) ([]*cloud_go_sdk.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServices", clusterID, limit, skip)
	ret0, _ := ret[0].([]*cloud_go_sdk.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServices indicates an expected call of ListServices.
func (mr *MockAPIMockRecorder) ListServices(clusterID, limit, skip interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServices", reflect.TypeOf((*MockAPI)(nil).ListServices), clusterID, limit, skip)
}

// Me mocks base method.
func (m *MockAPI) Me() (*cloud_go_sdk.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Me")
	ret0, _ := ret[0].(*cloud_go_sdk.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Me indicates an expected call of Me.
func (mr *MockAPIMockRecorder) Me() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Me", reflect.TypeOf((*MockAPI)(nil).Me))
}

// UpdateConsumer mocks base method.
func (m *MockAPI) UpdateConsumer(clusterID cloud_go_sdk.ID, consumer *cloud_go_sdk.Consumer) (*cloud_go_sdk.Consumer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateConsumer", clusterID, consumer)
	ret0, _ := ret[0].(*cloud_go_sdk.Consumer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateConsumer indicates an expected call of UpdateConsumer.
func (mr *MockAPIMockRecorder) UpdateConsumer(clusterID, consumer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateConsumer", reflect.TypeOf((*MockAPI)(nil).UpdateConsumer), clusterID, consumer)
}

// UpdateSSL mocks base method.
func (m *MockAPI) UpdateSSL(clusterID cloud_go_sdk.ID, ssl *cloud_go_sdk.Certificate) (*cloud_go_sdk.CertificateDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSSL", clusterID, ssl)
	ret0, _ := ret[0].(*cloud_go_sdk.CertificateDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSSL indicates an expected call of UpdateSSL.
func (mr *MockAPIMockRecorder) UpdateSSL(clusterID, ssl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSSL", reflect.TypeOf((*MockAPI)(nil).UpdateSSL), clusterID, ssl)
}

// UpdateService mocks base method.
func (m *MockAPI) UpdateService(clusterID cloud_go_sdk.ID, svc *cloud_go_sdk.Application) (*cloud_go_sdk.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateService", clusterID, svc)
	ret0, _ := ret[0].(*cloud_go_sdk.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateService indicates an expected call of UpdateService.
func (mr *MockAPIMockRecorder) UpdateService(clusterID, svc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateService", reflect.TypeOf((*MockAPI)(nil).UpdateService), clusterID, svc)
}
