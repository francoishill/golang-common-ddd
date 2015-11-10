package DbStorage

type DbStorageAndTransactionShared interface {
	Select(dest interface{}, query string, args ...interface{}) error
	Get(dest interface{}, query string, args ...interface{}) error

	Query(query string, args ...interface{}) (DbStorageScannableResultMultipleRows, error)
	QueryRow(query string, args ...interface{}) DbStorageScannableResultSingleRow

	Exec(query string, args ...interface{}) (DbStorageExecResult, error)
	MustExec(query string, args ...interface{}) DbStorageExecResult
}
