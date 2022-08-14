package accountpresenttest

import (
	"errors"
	"testing"

	accountpresent "github.com/ricky7171/te-marketplace/internal/modules/account/presentation"

	mocks "github.com/ricky7171/te-marketplace/internal/mocks/library_wrapper"
	mocksauthenticationservice "github.com/ricky7171/te-marketplace/internal/mocks/modules/account/application/service"

	"github.com/stretchr/testify/mock"
)

type returns []interface{}
type methodMock map[string]struct {
	params  params
	returns returns
}

type params []interface{}
type methodAssert map[string]params

func TestHandleLogin(t *testing.T) {

	// define testcase
	type input struct {
		ctxMock                   methodMock
		authenticationServiceMock methodMock
	}

	type expectedMethodCalls struct {
		ctxMock                   methodAssert
		authenticationServiceMock methodAssert
	}

	tests := []struct {
		name                string
		input               input
		expectedMethodCalls expectedMethodCalls
		expectedReturn      []interface{}
	}{
		{
			name: "Test success login",
			input: input{
				ctxMock: methodMock{
					"ShouldBindJSON": {
						params:  params{mock.Anything},
						returns: returns{nil},
					},
					"JSON": {
						params:  params{mock.Anything, mock.Anything},
						returns: returns{nil},
					},
				},
				authenticationServiceMock: methodMock{
					"Login": {
						params:  params{mock.Anything},
						returns: returns{true, nil},
					},
				},
			},
			expectedMethodCalls: expectedMethodCalls{
				ctxMock: methodAssert{
					"JSON": params{200, mock.Anything},
				},
				authenticationServiceMock: methodAssert{
					"Login": params{mock.Anything},
				},
			},
		},
		{
			name: "Test invalid request",
			input: input{
				ctxMock: methodMock{
					"ShouldBindJSON": {
						params:  params{mock.Anything},
						returns: returns{errors.New("invalid request")},
					},
					"JSON": {
						params:  params{mock.Anything, mock.Anything},
						returns: returns{nil},
					},
				},
			},
			expectedMethodCalls: expectedMethodCalls{
				ctxMock: methodAssert{
					"JSON": params{400, mock.Anything},
				},
			},
		},
	}

	// run testing
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// define object
			authenticationServiceMock := mocksauthenticationservice.NewAuthenticationService(t)
			handler := accountpresent.NewHandler(authenticationServiceMock)
			ctxMock := mocks.NewMyGinContext(t)

			// mock authentication service
			if test.input.authenticationServiceMock != nil {
				for methodName, methodMock := range test.input.authenticationServiceMock {
					authenticationServiceMock.Mock.On(methodName, methodMock.params...).Return(methodMock.returns...)
				}
			}

			// mock gin context
			if test.input.ctxMock != nil {
				for methodName, methodMock := range test.input.ctxMock {
					ctxMock.Mock.On(methodName, methodMock.params...).Return(methodMock.returns...)
				}
			}

			// do action
			handler.HandleLogin(ctxMock)

			// assert
			for key, val := range test.expectedMethodCalls.ctxMock {
				ctxMock.AssertCalled(t, key, val...)
			}

		})
	}

}
