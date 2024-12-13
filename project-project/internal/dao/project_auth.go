package dao

import (
	"context"
	"test.com/project-project/internal/data"
	"test.com/project-project/internal/database/gorms"
)

type ProjectAuthDao struct {
	conn *gorms.GormConn
}

func (p ProjectAuthDao) FindAuthList(ctx context.Context, orgCode int64) (list []*data.ProjectAuth, err error) {
	session := p.conn.Session(ctx)
	err = session.Model(&data.ProjectAuth{}).Where("organization_code=?", orgCode).Find(&list).Error
	return
}

func NewProjectAuthDao() *ProjectAuthDao {
	return &ProjectAuthDao{
		conn: gorms.New(),
	}
}
