package mysql

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		FindByMobile(ctx context.Context, mobile string) (*User, error)
		GetMaxUid() (int64, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

func (m *customUserModel) FindByMobile(ctx context.Context, mobile string) (*User, error) {
	return m.FindOneByMobile(ctx, mobile)
}

func (m *customUserModel) GetMaxUid() (int64, error) {
	query := "select max(uid) from user"
	var maxUid int64
	err := m.conn.QueryRow(maxUid, query)
	if err != nil {
		return 0, err
	}
	return maxUid, nil
}

// NewUserModel returns a model for the database table.
func NewUserModel(dsn string) UserModel {
	conn := sqlx.NewMysql(dsn)
	return &customUserModel{
		defaultUserModel: newUserModel(conn),
	}
}
