package accountappservice

import (
	"errors"

	"github.com/ricky7171/te-marketplace/internal/helper"

	accountdom "github.com/ricky7171/te-marketplace/internal/modules/account/domain"
	accountdomrepository "github.com/ricky7171/te-marketplace/internal/modules/account/domain/repository"
)

type AuthenticationService interface {
	Login(email string, password string) (string, string, *accountdom.Account, error)
}

// implementation
type AuthenticationServiceImpl struct {
	accountRepository accountdomrepository.AccountRepository
	jwtHelper         helper.HelperJwt
}

func NewAuthenticationServiceImpl(accountRepository accountdomrepository.AccountRepository, jwtHelper helper.HelperJwt) *AuthenticationServiceImpl {
	return &AuthenticationServiceImpl{
		accountRepository: accountRepository,
		jwtHelper:         jwtHelper,
	}
}

func (service *AuthenticationServiceImpl) Login(email string, password string) (string, string, *accountdom.Account, error) {
	// get account by email and password
	account, err := service.accountRepository.GetByFields(map[string]string{
		"email":    email,
		"password": password,
	})

	if err != nil {
		return "", "", nil, errors.New("email / password not found")
	}

	//generate JWT
	token, refreshToken, err := service.jwtHelper.GenerateToken(account.Email, *account.Id)
	if err != nil {
		return "", "", nil, errors.New("failed to generate token & refresh token")
	}

	return token, refreshToken, account, nil

}
