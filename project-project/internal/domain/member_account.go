package domain

import (
	"test.com/project-common/errs"
	"test.com/project-project/internal/dao"
	"test.com/project-project/internal/data"
	"test.com/project-project/internal/repo"
)

type AccountDomain struct {
	accountRepo   repo.AccountRepo
	userRpcDomain *UserRpcDomain
}

func (d AccountDomain) AccountList(organizationCode string,
	memberId int64,
	page int64,
	pageSize int64,
	departmentCode string,
	searchType int32) ([]*data.MemberAccountDisplay, int64, *errs.BError) {
	return nil, 0, nil
}

func NewAccountDomain() *AccountDomain {
	return &AccountDomain{
		accountRepo:   dao.NewTaskWorkTimeDao(),
		userRpcDomain: NewUserRpcDomain(),
	}
}
