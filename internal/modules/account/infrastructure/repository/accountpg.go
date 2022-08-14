package accountinfrarepository

import (
	accountdom "github.com/ricky7171/te-marketplace/internal/modules/account/domain"
)

type AccountRepositoryPg struct {
}

func NewAccountRepositoryPg() *AccountRepositoryPg {
	return &AccountRepositoryPg{}
}

func (repo *AccountRepositoryPg) GetByFields(account accountdom.Account, fields []string) (interface{}, error) {
	return nil, nil
}
