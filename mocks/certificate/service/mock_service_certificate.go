// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_certificateService is a generated GoMock package.
package mock_certificateService

import (
	reflect "reflect"
	certificateEntities "sirclo/project-capstone/entities/certificateEntities"
	certificateRequest "sirclo/project-capstone/utils/request/certificateRequest"

	gomock "github.com/golang/mock/gomock"
)

// MockCertificateServiceInterface is a mock of CertificateServiceInterface interface.
type MockCertificateServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCertificateServiceInterfaceMockRecorder
}

// MockCertificateServiceInterfaceMockRecorder is the mock recorder for MockCertificateServiceInterface.
type MockCertificateServiceInterfaceMockRecorder struct {
	mock *MockCertificateServiceInterface
}

// NewMockCertificateServiceInterface creates a new mock instance.
func NewMockCertificateServiceInterface(ctrl *gomock.Controller) *MockCertificateServiceInterface {
	mock := &MockCertificateServiceInterface{ctrl: ctrl}
	mock.recorder = &MockCertificateServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCertificateServiceInterface) EXPECT() *MockCertificateServiceInterfaceMockRecorder {
	return m.recorder
}

// CountVaccineIsPending mocks base method.
func (m *MockCertificateServiceInterface) CountVaccineIsPending(userID string, dossage int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountVaccineIsPending", userID, dossage)
	ret0, _ := ret[0].(int)
	return ret0
}

// CountVaccineIsPending indicates an expected call of CountVaccineIsPending.
func (mr *MockCertificateServiceInterfaceMockRecorder) CountVaccineIsPending(userID, dossage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountVaccineIsPending", reflect.TypeOf((*MockCertificateServiceInterface)(nil).CountVaccineIsPending), userID, dossage)
}

// GetCertificate mocks base method.
func (m *MockCertificateServiceInterface) GetCertificate(id string) (certificateEntities.Certificate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCertificate", id)
	ret0, _ := ret[0].(certificateEntities.Certificate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCertificate indicates an expected call of GetCertificate.
func (mr *MockCertificateServiceInterfaceMockRecorder) GetCertificate(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCertificate", reflect.TypeOf((*MockCertificateServiceInterface)(nil).GetCertificate), id)
}

// GetCertificateUser mocks base method.
func (m *MockCertificateServiceInterface) GetCertificateUser(userId string) ([]certificateEntities.Certificate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCertificateUser", userId)
	ret0, _ := ret[0].([]certificateEntities.Certificate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCertificateUser indicates an expected call of GetCertificateUser.
func (mr *MockCertificateServiceInterfaceMockRecorder) GetCertificateUser(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCertificateUser", reflect.TypeOf((*MockCertificateServiceInterface)(nil).GetCertificateUser), userId)
}

// GetCertificates mocks base method.
func (m *MockCertificateServiceInterface) GetCertificates(orderBy string) ([]certificateEntities.Certificates, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCertificates", orderBy)
	ret0, _ := ret[0].([]certificateEntities.Certificates)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCertificates indicates an expected call of GetCertificates.
func (mr *MockCertificateServiceInterfaceMockRecorder) GetCertificates(orderBy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCertificates", reflect.TypeOf((*MockCertificateServiceInterface)(nil).GetCertificates), orderBy)
}

// GetVaccineDose mocks base method.
func (m *MockCertificateServiceInterface) GetVaccineDose(userID, status string) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVaccineDose", userID, status)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetVaccineDose indicates an expected call of GetVaccineDose.
func (mr *MockCertificateServiceInterfaceMockRecorder) GetVaccineDose(userID, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVaccineDose", reflect.TypeOf((*MockCertificateServiceInterface)(nil).GetVaccineDose), userID, status)
}

// UploadCertificateVaccine mocks base method.
func (m *MockCertificateServiceInterface) UploadCertificateVaccine(userID string, input certificateRequest.CertificateUploadRequest) (certificateEntities.Certificate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadCertificateVaccine", userID, input)
	ret0, _ := ret[0].(certificateEntities.Certificate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadCertificateVaccine indicates an expected call of UploadCertificateVaccine.
func (mr *MockCertificateServiceInterfaceMockRecorder) UploadCertificateVaccine(userID, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadCertificateVaccine", reflect.TypeOf((*MockCertificateServiceInterface)(nil).UploadCertificateVaccine), userID, input)
}

// VerifyCertificate mocks base method.
func (m *MockCertificateServiceInterface) VerifyCertificate(id, userID string, input certificateRequest.CertificateVerificationRequest) (certificateEntities.Certificate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyCertificate", id, userID, input)
	ret0, _ := ret[0].(certificateEntities.Certificate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyCertificate indicates an expected call of VerifyCertificate.
func (mr *MockCertificateServiceInterfaceMockRecorder) VerifyCertificate(id, userID, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyCertificate", reflect.TypeOf((*MockCertificateServiceInterface)(nil).VerifyCertificate), id, userID, input)
}