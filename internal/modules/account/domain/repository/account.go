package accountdomrepository

import (
	accountdom "github.com/ricky7171/te-marketplace/internal/modules/account/domain"
)

type AccountRepository interface {
	GetByFields(fields map[string]string) (*accountdom.Account, error)
}
