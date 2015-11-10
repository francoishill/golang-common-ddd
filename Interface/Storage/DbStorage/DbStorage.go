package DbStorage

type DbStorage interface {
	DbStorageAndTransactionShared

	Migrate() (numMigrationsApplied int)

	BeginTransaction() (DbStorageTransaction, error)
	MustBeginTransaction() DbStorageTransaction

	DeferableCommitOnSuccessRollbackOnFail(tx DbStorageTransaction)

	IsNoRowsError(err error) bool
}
