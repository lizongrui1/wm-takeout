package tx

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var TRANSACTIONS_DB_CANNOTRESOLVE = errors.New("事务接口无法转换为具体实现类")

type GormTransaction struct {
	db *gorm.DB
}

func NewGormTransaction(db *gorm.DB, ctx context.Context) Transaction {
	return &GormTransaction{
		db: db.WithContext(ctx),
	}
}

func (g *GormTransaction) Begin() error {
	g.db = g.db.Begin()
	return g.db.Error
}

func (g *GormTransaction) Commit() error {
	g.db = g.db.Commit()
	return g.db.Error
}

func (g *GormTransaction) Rollback() {
	g.db.Rollback()
}

func GetGormDB(transactions Transaction) (*gorm.DB, error) {
	if gormTx, ok := transactions.(*GormTransaction); ok {
		return gormTx.db, nil
	}
	return nil, TRANSACTIONS_DB_CANNOTRESOLVE
}
