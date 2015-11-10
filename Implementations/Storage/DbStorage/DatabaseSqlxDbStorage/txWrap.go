package DatabaseSqlxDbStorage

import (
	"github.com/jmoiron/sqlx"

	. "github.com/francoishill/golang-common-ddd/Interface/Storage/DbStorage"
)

type txWrap struct {
	tx *sqlx.Tx
}

func (t *txWrap) Select(dest interface{}, query string, args ...interface{}) error {
	return t.Select(dest, query, args...)
}
func (t *txWrap) Get(dest interface{}, query string, args ...interface{}) error {
	return t.Get(dest, query, args...)
}
func (t *txWrap) Query(query string, args ...interface{}) (DbStorageScannableResultMultipleRows, error) {
	return t.Query(query, args...)
}
func (t *txWrap) QueryRow(query string, args ...interface{}) DbStorageScannableResultSingleRow {
	return t.QueryRow(query, args...)
}
func (t *txWrap) Exec(query string, args ...interface{}) (DbStorageExecResult, error) {
	return t.Exec(query, args...)
}
func (t *txWrap) MustExec(query string, args ...interface{}) DbStorageExecResult {
	return t.MustExec(query, args...)
}
func (t *txWrap) Commit() error {
	return t.Commit()
}
func (t *txWrap) Rollback() error {
	return t.Rollback()
}
