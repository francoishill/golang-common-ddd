package DbStorage

type DbStorageExecResult interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}
