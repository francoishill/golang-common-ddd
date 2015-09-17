package DbStorage

type DbStorageTransaction interface {
	Commit() error
	Rollback() error
}
