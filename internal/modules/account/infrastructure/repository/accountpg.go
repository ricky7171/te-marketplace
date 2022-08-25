package accountinfrarepository

import (
	accountdom "github.com/ricky7171/te-marketplace/internal/modules/account/domain"
)

type AccountRepositoryPg struct {
}

func NewAccountRepositoryPg() *AccountRepositoryPg {
	return &AccountRepositoryPg{}
}

func (repo *AccountRepositoryPg) GetByFields(fields map[string]string) (*accountdom.Account, error) {
	return nil, nil
}
