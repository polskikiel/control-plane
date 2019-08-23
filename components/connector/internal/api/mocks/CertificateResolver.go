// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import externalschema "github.com/kyma-incubator/compass/components/connector/pkg/graphql/externalschema"
import mock "github.com/stretchr/testify/mock"

// CertificateResolver is an autogenerated mock type for the CertificateResolver type
type CertificateResolver struct {
	mock.Mock
}

// Configuration provides a mock function with given fields: ctx
func (_m *CertificateResolver) Configuration(ctx context.Context) (*externalschema.Configuration, error) {
	ret := _m.Called(ctx)

	var r0 *externalschema.Configuration
	if rf, ok := ret.Get(0).(func(context.Context) *externalschema.Configuration); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*externalschema.Configuration)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RevokeCertificate provides a mock function with given fields: ctx
func (_m *CertificateResolver) RevokeCertificate(ctx context.Context) (bool, error) {
	ret := _m.Called(ctx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignCertificateSigningRequest provides a mock function with given fields: ctx, csr
func (_m *CertificateResolver) SignCertificateSigningRequest(ctx context.Context, csr string) (*externalschema.CertificationResult, error) {
	ret := _m.Called(ctx, csr)

	var r0 *externalschema.CertificationResult
	if rf, ok := ret.Get(0).(func(context.Context, string) *externalschema.CertificationResult); ok {
		r0 = rf(ctx, csr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*externalschema.CertificationResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, csr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}