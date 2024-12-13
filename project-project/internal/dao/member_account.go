package dao

import "test.com/project-project/internal/database/gorms"

type MemberAccountDao struct {
	conn *gorms.GormConn
}

func NewMemberAccountDao() *MemberAccountDao {
	return &MemberAccountDao{
		conn: gorms.New(),
	}
}
