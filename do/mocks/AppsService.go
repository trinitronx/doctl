// Code generated by MockGen. DO NOT EDIT.
// Source: apps.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	godo "github.com/digitalocean/godo"
	gomock "github.com/golang/mock/gomock"
)

// MockAppsService is a mock of AppsService interface.
type MockAppsService struct {
	ctrl     *gomock.Controller
	recorder *MockAppsServiceMockRecorder
}

// MockAppsServiceMockRecorder is the mock recorder for MockAppsService.
type MockAppsServiceMockRecorder struct {
	mock *MockAppsService
}

// NewMockAppsService creates a new mock instance.
func NewMockAppsService(ctrl *gomock.Controller) *MockAppsService {
	mock := &MockAppsService{ctrl: ctrl}
	mock.recorder = &MockAppsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppsService) EXPECT() *MockAppsServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAppsService) Create(req *godo.AppCreateRequest) (*godo.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", req)
	ret0, _ := ret[0].(*godo.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAppsServiceMockRecorder) Create(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAppsService)(nil).Create), req)
}

// CreateDeployment mocks base method.
func (m *MockAppsService) CreateDeployment(appID string, forceRebuild bool) (*godo.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDeployment", appID, forceRebuild)
	ret0, _ := ret[0].(*godo.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDeployment indicates an expected call of CreateDeployment.
func (mr *MockAppsServiceMockRecorder) CreateDeployment(appID, forceRebuild interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDeployment", reflect.TypeOf((*MockAppsService)(nil).CreateDeployment), appID, forceRebuild)
}

// Delete mocks base method.
func (m *MockAppsService) Delete(appID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", appID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAppsServiceMockRecorder) Delete(appID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAppsService)(nil).Delete), appID)
}

// Get mocks base method.
func (m *MockAppsService) Get(appID string) (*godo.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", appID)
	ret0, _ := ret[0].(*godo.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockAppsServiceMockRecorder) Get(appID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAppsService)(nil).Get), appID)
}

// GetDeployment mocks base method.
func (m *MockAppsService) GetDeployment(appID, deploymentID string) (*godo.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeployment", appID, deploymentID)
	ret0, _ := ret[0].(*godo.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeployment indicates an expected call of GetDeployment.
func (mr *MockAppsServiceMockRecorder) GetDeployment(appID, deploymentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeployment", reflect.TypeOf((*MockAppsService)(nil).GetDeployment), appID, deploymentID)
}

// GetInstanceSize mocks base method.
func (m *MockAppsService) GetInstanceSize(slug string) (*godo.AppInstanceSize, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstanceSize", slug)
	ret0, _ := ret[0].(*godo.AppInstanceSize)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstanceSize indicates an expected call of GetInstanceSize.
func (mr *MockAppsServiceMockRecorder) GetInstanceSize(slug interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstanceSize", reflect.TypeOf((*MockAppsService)(nil).GetInstanceSize), slug)
}

// GetLogs mocks base method.
func (m *MockAppsService) GetLogs(appID, deploymentID, component string, logType godo.AppLogType, follow bool, tail int) (*godo.AppLogs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogs", appID, deploymentID, component, logType, follow, tail)
	ret0, _ := ret[0].(*godo.AppLogs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLogs indicates an expected call of GetLogs.
func (mr *MockAppsServiceMockRecorder) GetLogs(appID, deploymentID, component, logType, follow, tail interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogs", reflect.TypeOf((*MockAppsService)(nil).GetLogs), appID, deploymentID, component, logType, follow, tail)
}

// GetTier mocks base method.
func (m *MockAppsService) GetTier(slug string) (*godo.AppTier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTier", slug)
	ret0, _ := ret[0].(*godo.AppTier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTier indicates an expected call of GetTier.
func (mr *MockAppsServiceMockRecorder) GetTier(slug interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTier", reflect.TypeOf((*MockAppsService)(nil).GetTier), slug)
}

// List mocks base method.
func (m *MockAppsService) List() ([]*godo.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*godo.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockAppsServiceMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockAppsService)(nil).List))
}

// ListAlerts mocks base method.
func (m *MockAppsService) ListAlerts(appID string) ([]*godo.AppAlert, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAlerts", appID)
	ret0, _ := ret[0].([]*godo.AppAlert)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAlerts indicates an expected call of ListAlerts.
func (mr *MockAppsServiceMockRecorder) ListAlerts(appID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAlerts", reflect.TypeOf((*MockAppsService)(nil).ListAlerts), appID)
}

// ListDeployments mocks base method.
func (m *MockAppsService) ListDeployments(appID string) ([]*godo.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDeployments", appID)
	ret0, _ := ret[0].([]*godo.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDeployments indicates an expected call of ListDeployments.
func (mr *MockAppsServiceMockRecorder) ListDeployments(appID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDeployments", reflect.TypeOf((*MockAppsService)(nil).ListDeployments), appID)
}

// ListInstanceSizes mocks base method.
func (m *MockAppsService) ListInstanceSizes() ([]*godo.AppInstanceSize, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListInstanceSizes")
	ret0, _ := ret[0].([]*godo.AppInstanceSize)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInstanceSizes indicates an expected call of ListInstanceSizes.
func (mr *MockAppsServiceMockRecorder) ListInstanceSizes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInstanceSizes", reflect.TypeOf((*MockAppsService)(nil).ListInstanceSizes))
}

// ListRegions mocks base method.
func (m *MockAppsService) ListRegions() ([]*godo.AppRegion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRegions")
	ret0, _ := ret[0].([]*godo.AppRegion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRegions indicates an expected call of ListRegions.
func (mr *MockAppsServiceMockRecorder) ListRegions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRegions", reflect.TypeOf((*MockAppsService)(nil).ListRegions))
}

// ListTiers mocks base method.
func (m *MockAppsService) ListTiers() ([]*godo.AppTier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTiers")
	ret0, _ := ret[0].([]*godo.AppTier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTiers indicates an expected call of ListTiers.
func (mr *MockAppsServiceMockRecorder) ListTiers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTiers", reflect.TypeOf((*MockAppsService)(nil).ListTiers))
}

// Propose mocks base method.
func (m *MockAppsService) Propose(req *godo.AppProposeRequest) (*godo.AppProposeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Propose", req)
	ret0, _ := ret[0].(*godo.AppProposeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Propose indicates an expected call of Propose.
func (mr *MockAppsServiceMockRecorder) Propose(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Propose", reflect.TypeOf((*MockAppsService)(nil).Propose), req)
}

// Update mocks base method.
func (m *MockAppsService) Update(appID string, req *godo.AppUpdateRequest) (*godo.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", appID, req)
	ret0, _ := ret[0].(*godo.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockAppsServiceMockRecorder) Update(appID, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAppsService)(nil).Update), appID, req)
}

// UpdateAlertDestinations mocks base method.
func (m *MockAppsService) UpdateAlertDestinations(appID, alertID string, update *godo.AlertDestinationUpdateRequest) (*godo.AppAlert, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAlertDestinations", appID, alertID, update)
	ret0, _ := ret[0].(*godo.AppAlert)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAlertDestinations indicates an expected call of UpdateAlertDestinations.
func (mr *MockAppsServiceMockRecorder) UpdateAlertDestinations(appID, alertID, update interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAlertDestinations", reflect.TypeOf((*MockAppsService)(nil).UpdateAlertDestinations), appID, alertID, update)
}
