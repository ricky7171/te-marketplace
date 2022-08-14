package accountdomrepository

import (
	accountdom "github.com/ricky7171/te-marketplace/internal/modules/account/domain"
)

type AccountRepository interface {
	GetByFields(account accountdom.Account, fields []string) (interface{}, error)
}
