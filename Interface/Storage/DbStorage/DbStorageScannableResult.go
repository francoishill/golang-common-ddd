package DbStorage

type DbStorageScannableResultSingleRow interface {
	Columns() ([]string, error)
	MapScan(dest map[string]interface{}) error
	Scan(dest ...interface{}) error
	StructScan(dest interface{}) error
}

type DbStorageScannableResultMultipleRows interface {
	DbStorageScannableResultSingleRow
	Next() bool
	Close() error
}
