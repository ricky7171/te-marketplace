package accountappservice

import (
	accountdom "github.com/ricky7171/te-marketplace/internal/modules/account/domain"
)

type AuthenticationService interface {
	Login(loginRequest accountdom.Credential) (interface{}, error)
}

// implementation

type AuthenticationServiceImpl struct{}

func NewAuthenticationServiceImpl() *AuthenticationServiceImpl {
	return &AuthenticationServiceImpl{}
}

func (service *AuthenticationServiceImpl) Login(loginRequest accountdom.Credential) (interface{}, error) {
	return nil, nil
}
