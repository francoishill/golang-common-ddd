package DbStorage

type DbStorageTransaction interface {
	DbStorageAndTransactionShared

	Commit() error
	Rollback() error
}
