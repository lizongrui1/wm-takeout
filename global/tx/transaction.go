package tx

type Transaction interface {
	// Begin 开始一个事务
	Begin() error
	// Commit 提交事务
	Commit() error
	// Rollback 回滚事务
	Rollback()
}
