package helpertest

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ricky7171/te-marketplace/internal/helper"

	"github.com/golang-jwt/jwt/v4"

	mocks "github.com/ricky7171/te-marketplace/internal/mocks/library_wrapper"
	"github.com/stretchr/testify/mock"
)

type returns []interface{}
type methodMock map[string]struct {
	params  params
	returns returns
}

type params []interface{}
type methodAssert map[string]any

func TestGenerateToken(t *testing.T) {

	// define testcase
	type input struct {
		myJwtMock methodMock
		name      string
		userId    int
	}

	type expectedMethodCalls struct {
		myJwtMock methodAssert
	}

	tests := []struct {
		name                string
		input               input
		expectedMethodCalls expectedMethodCalls
		expectedReturn      []interface{}
	}{
		{
			name: "Test success generate token",
			input: input{
				myJwtMock: methodMock{
					"GenerateStandardClaims": {
						params: params{mock.Anything},
						returns: returns{jwt.RegisteredClaims{
							ExpiresAt: &jwt.NumericDate{
								Time: time.Now().Local().Add(time.Hour * time.Duration(16)),
							},
						}},
					},
					"NewToken": {
						params:  params{mock.Anything, mock.Anything},
						returns: returns{"token123", nil},
					},
				},
				name:   "Ricky",
				userId: 12,
			},
			expectedMethodCalls: expectedMethodCalls{
				myJwtMock: methodAssert{
					"GenerateStandardClaims": []params{
						params{24},
						params{168},
					},
					"NewToken": 2,
				},
			},
			expectedReturn: []any{"token123", "token123", nil},
		},
		{
			name: "Test fail generate new token",
			input: input{
				myJwtMock: methodMock{
					"GenerateStandardClaims": {
						params: params{mock.Anything},
						returns: returns{jwt.RegisteredClaims{
							ExpiresAt: &jwt.NumericDate{
								Time: time.Now().Local().Add(time.Hour * time.Duration(16)),
							},
						}},
					},
					"NewToken": {
						params:  params{mock.Anything, mock.Anything},
						returns: returns{"", errors.New("failed to generate new token")},
					},
				},
				name:   "Ricky",
				userId: 12,
			},
			expectedMethodCalls: expectedMethodCalls{
				myJwtMock: methodAssert{
					"GenerateStandardClaims": []params{
						params{24},
						params{168},
					},
					"NewToken": 1,
				},
			},
			expectedReturn: []any{"", "", errors.New("failed to generate new token")},
		},
		{
			name: "Test invalid parameter",
			input: input{
				myJwtMock: methodMock{},
				name:      "",
				userId:    0,
			},
			expectedMethodCalls: expectedMethodCalls{},
			expectedReturn:      []any{"", "", errors.New("name or userid cannot be empty")},
		},
	}

	// run testing
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// define object
			myJwtMock := mocks.NewMyJwt(t)
			jwtHelper := helper.NewHelperJwt(myJwtMock)

			// mock myjwt
			if test.input.myJwtMock != nil {
				for methodName, methodMock := range test.input.myJwtMock {
					myJwtMock.Mock.On(methodName, methodMock.params...).Return(methodMock.returns...)
				}
			}

			// do action
			token, refreshToken, err := jwtHelper.GenerateToken(test.input.name, test.input.userId)

			// assert calling
			for methodName, assert := range test.expectedMethodCalls.myJwtMock {
				typeOfAssert := reflect.TypeOf(assert).String()
				if typeOfAssert == "int" { //it means, we need to assert count of call
					myJwtMock.AssertNumberOfCalls(t, methodName, assert.(int))
				} else if typeOfAssert == "[]helpertest.params" { //if []helpertest.params, then assert of calling method with specific parameter
					for _, callAssert := range assert.([]params) {
						myJwtMock.AssertCalled(t, methodName, callAssert...)
					}
				}
			}

			// assert return
			assert.Equal(t, test.expectedReturn[0], token)
			assert.Equal(t, test.expectedReturn[1], refreshToken)
			assert.Equal(t, test.expectedReturn[2], err)

		})
	}

}

func TestValidateToken(t *testing.T) {

	// define testcase
	type input struct {
		myJwtMock   methodMock
		signedToken string
	}

	type expectedMethodCalls struct {
		myJwtMock methodAssert
	}

	// temporary
	time16HourAfterNow := time.Now().Local().Add(time.Hour * 16)
	time16HourBeforeNow := time.Now().Local().Add(time.Hour * -16)

	tests := []struct {
		name                string
		input               input
		expectedMethodCalls expectedMethodCalls
		expectedReturn      []interface{}
	}{
		{
			name: "Test success validate token that expired in 16h",
			input: input{
				myJwtMock: methodMock{
					"ParseWithClaims": {
						params: params{mock.Anything, mock.Anything, mock.Anything, mock.Anything},
						returns: returns{
							&jwt.Token{
								Claims: &helper.SignedTokenDetails{
									RegisteredClaims: jwt.RegisteredClaims{
										ExpiresAt: &jwt.NumericDate{
											Time: time16HourAfterNow,
										},
									},
									Name: "ricky",
									ID:   "123",
								},
							}, nil,
						},
					},
				},
				signedToken: "token123",
			},
			expectedMethodCalls: expectedMethodCalls{
				myJwtMock: methodAssert{
					"ParseWithClaims": 1,
				},
			},
			expectedReturn: []any{&helper.SignedTokenDetails{
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: &jwt.NumericDate{
						Time: time16HourAfterNow,
					},
				},
				Name: "ricky",
				ID:   "123",
			}, nil},
		},
		{
			name: "Test with token that already expired",
			input: input{
				myJwtMock: methodMock{
					"ParseWithClaims": {
						params: params{mock.Anything, mock.Anything, mock.Anything, mock.Anything},
						returns: returns{
							&jwt.Token{
								Claims: &helper.SignedTokenDetails{
									RegisteredClaims: jwt.RegisteredClaims{
										ExpiresAt: &jwt.NumericDate{
											Time: time16HourBeforeNow,
										},
									},
									Name: "ricky",
									ID:   "123",
								},
							}, nil,
						},
					},
				},
				signedToken: "token123",
			},
			expectedMethodCalls: expectedMethodCalls{
				myJwtMock: methodAssert{
					"ParseWithClaims": 1,
				},
			},
			expectedReturn: []any{nil, errors.New("the token is expired")},
		},
		{
			name: "Test fail parsewithclaims",
			input: input{
				myJwtMock: methodMock{
					"ParseWithClaims": {
						params: params{mock.Anything, mock.Anything, mock.Anything, mock.Anything},
						returns: returns{
							nil, errors.New("failed to parse"),
						},
					},
				},
				signedToken: "token123",
			},
			expectedMethodCalls: expectedMethodCalls{
				myJwtMock: methodAssert{
					"ParseWithClaims": 1,
				},
			},
			expectedReturn: []any{nil, errors.New("failed to parse")},
		},
		{
			name: "Test empty input",
			input: input{
				myJwtMock:   methodMock{},
				signedToken: "",
			},
			expectedMethodCalls: expectedMethodCalls{
				myJwtMock: methodAssert{
					"ParseWithClaims": 0,
				},
			},
			expectedReturn: []any{nil, errors.New("empty signed token")},
		},
	}

	// run testing
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// define object
			myJwtMock := mocks.NewMyJwt(t)
			jwtHelper := helper.NewHelperJwt(myJwtMock)

			// mock myjwt
			if test.input.myJwtMock != nil {
				for methodName, methodMock := range test.input.myJwtMock {
					myJwtMock.Mock.On(methodName, methodMock.params...).Return(methodMock.returns...)
				}
			}

			// do action
			fmt.Println("check test.input.signedtoken")
			fmt.Println(test.input.signedToken)
			claims, err := jwtHelper.ValidateToken(test.input.signedToken)

			// assert calling
			for methodName, assert := range test.expectedMethodCalls.myJwtMock {
				typeOfAssert := reflect.TypeOf(assert).String()
				if typeOfAssert == "int" { //it means, we need to assert count of call
					myJwtMock.AssertNumberOfCalls(t, methodName, assert.(int))
				} else if typeOfAssert == "[]helpertest.params" { //if []helpertest.params, then assert of calling method with specific parameter
					for _, callAssert := range assert.([]params) {
						myJwtMock.AssertCalled(t, methodName, callAssert...)
					}
				}
			}

			// assert return
			if test.expectedReturn[0] == nil {
				assert.Nil(t, claims)
			} else {
				assert.Equal(t, test.expectedReturn[0], claims)
			}

			if test.expectedReturn[1] == nil {
				assert.Nil(t, err)
			} else {
				assert.Equal(t, test.expectedReturn[1], err)
			}

			if claims == nil {
				fmt.Println("claim adalah nil")
				fmt.Println(claims)
				fmt.Println(test.expectedReturn[0])
			} else {
				fmt.Println("claim tidak nil")
				fmt.Println(claims)
			}

		})
	}

}

func TestValidateRefreshToken(t *testing.T) {

	// define testcase
	type input struct {
		myJwtMock   methodMock
		signedToken string
	}

	type expectedMethodCalls struct {
		myJwtMock methodAssert
	}

	// temporary
	time16HourAfterNow := time.Now().Local().Add(time.Hour * 16)
	time16HourBeforeNow := time.Now().Local().Add(time.Hour * -16)

	tests := []struct {
		name                string
		input               input
		expectedMethodCalls expectedMethodCalls
		expectedReturn      []interface{}
	}{
		{
			name: "Test success validate refresh token that expired in 16h",
			input: input{
				myJwtMock: methodMock{
					"ParseWithClaims": {
						params: params{mock.Anything, mock.Anything, mock.Anything, mock.Anything},
						returns: returns{
							&jwt.Token{
								Claims: &helper.SignedRefreshTokenDetails{
									RegisteredClaims: jwt.RegisteredClaims{
										ExpiresAt: &jwt.NumericDate{
											Time: time16HourAfterNow,
										},
									},
									Name: "ricky",
									ID:   "123",
								},
							}, nil,
						},
					},
				},
				signedToken: "token123",
			},
			expectedMethodCalls: expectedMethodCalls{
				myJwtMock: methodAssert{
					"ParseWithClaims": 1,
				},
			},
			expectedReturn: []any{&helper.SignedRefreshTokenDetails{
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: &jwt.NumericDate{
						Time: time16HourAfterNow,
					},
				},
				Name: "ricky",
				ID:   "123",
			}, nil},
		},
		{
			name: "Test with refresh token that already expired",
			input: input{
				myJwtMock: methodMock{
					"ParseWithClaims": {
						params: params{mock.Anything, mock.Anything, mock.Anything, mock.Anything},
						returns: returns{
							&jwt.Token{
								Claims: &helper.SignedRefreshTokenDetails{
									RegisteredClaims: jwt.RegisteredClaims{
										ExpiresAt: &jwt.NumericDate{
											Time: time16HourBeforeNow,
										},
									},
									Name: "ricky",
									ID:   "123",
								},
							}, nil,
						},
					},
				},
				signedToken: "token123",
			},
			expectedMethodCalls: expectedMethodCalls{
				myJwtMock: methodAssert{
					"ParseWithClaims": 1,
				},
			},
			expectedReturn: []any{nil, errors.New("the refresh token is expired")},
		},
		{
			name: "Test fail parsewithclaims",
			input: input{
				myJwtMock: methodMock{
					"ParseWithClaims": {
						params: params{mock.Anything, mock.Anything, mock.Anything, mock.Anything},
						returns: returns{
							nil, errors.New("failed to parse"),
						},
					},
				},
				signedToken: "token123",
			},
			expectedMethodCalls: expectedMethodCalls{
				myJwtMock: methodAssert{
					"ParseWithClaims": 1,
				},
			},
			expectedReturn: []any{nil, errors.New("failed to parse")},
		},
		{
			name: "Test empty input",
			input: input{
				myJwtMock:   methodMock{},
				signedToken: "",
			},
			expectedMethodCalls: expectedMethodCalls{
				myJwtMock: methodAssert{
					"ParseWithClaims": 0,
				},
			},
			expectedReturn: []any{nil, errors.New("empty signed refresh token")},
		},
	}

	// run testing
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// define object
			myJwtMock := mocks.NewMyJwt(t)
			jwtHelper := helper.NewHelperJwt(myJwtMock)

			// mock myjwt
			if test.input.myJwtMock != nil {
				for methodName, methodMock := range test.input.myJwtMock {
					myJwtMock.Mock.On(methodName, methodMock.params...).Return(methodMock.returns...)
				}
			}

			// do action
			fmt.Println("check test.input.signedtoken")
			fmt.Println(test.input.signedToken)
			claims, err := jwtHelper.ValidateRefreshToken(test.input.signedToken)

			// assert calling
			for methodName, assert := range test.expectedMethodCalls.myJwtMock {
				typeOfAssert := reflect.TypeOf(assert).String()
				if typeOfAssert == "int" { //it means, we need to assert count of call
					myJwtMock.AssertNumberOfCalls(t, methodName, assert.(int))
				} else if typeOfAssert == "[]helpertest.params" { //if []helpertest.params, then assert of calling method with specific parameter
					for _, callAssert := range assert.([]params) {
						myJwtMock.AssertCalled(t, methodName, callAssert...)
					}
				}
			}

			// assert return
			if test.expectedReturn[0] == nil {
				assert.Nil(t, claims)
			} else {
				assert.Equal(t, test.expectedReturn[0], claims)
			}

			if test.expectedReturn[1] == nil {
				assert.Nil(t, err)
			} else {
				assert.Equal(t, test.expectedReturn[1], err)
			}

			if claims == nil {
				fmt.Println("claim adalah nil")
				fmt.Println(claims)
				fmt.Println(test.expectedReturn[0])
			} else {
				fmt.Println("claim tidak nil")
				fmt.Println(claims)
			}

		})
	}

}
