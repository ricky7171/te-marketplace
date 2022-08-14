package accountappservice

import (
	"errors"
	"fmt"

	accountdom "github.com/ricky7171/te-marketplace/internal/modules/account/domain"
	accountdomrepository "github.com/ricky7171/te-marketplace/internal/modules/account/domain/repository"
)

type AuthenticationService interface {
	Login(credential accountdom.Credential) (interface{}, error)
}

// implementation
type AuthenticationServiceImpl struct {
	accountRepository accountdomrepository.AccountRepository
}

func NewAuthenticationServiceImpl(accountRepository accountdomrepository.AccountRepository) *AuthenticationServiceImpl {
	return &AuthenticationServiceImpl{
		accountRepository: accountRepository,
	}
}

func (service *AuthenticationServiceImpl) Login(credential accountdom.Credential) (interface{}, error) {
	account := accountdom.NewAccount(nil, credential.Email, credential.Password, nil, nil, nil)
	result, err := service.accountRepository.GetByFields(*account, []string{"email", "password"})
	if err != nil {
		return nil, errors.New("email / password tidak ditemukan")
	}

	//generate JWT

	//return jwt and account data
	fmt.Println("check result")
	fmt.Println(result)

	return account, nil

}
