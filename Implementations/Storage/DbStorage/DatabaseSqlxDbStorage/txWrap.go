package DatabaseSqlxDbStorage

import (
	"github.com/jmoiron/sqlx"

	. "github.com/francoishill/golang-common-ddd/Interface/Storage/DbStorage"
)

type txWrap struct {
	tx *sqlx.Tx
}

func (t *txWrap) Select(dest interface{}, query string, args ...interface{}) error {
	return t.tx.Select(dest, query, args...)
}
func (t *txWrap) Get(dest interface{}, query string, args ...interface{}) error {
	return t.tx.Get(dest, query, args...)
}
func (t *txWrap) Query(query string, args ...interface{}) (DbStorageScannableResultMultipleRows, error) {
	return t.tx.Queryx(query, args...)
}
func (t *txWrap) QueryRow(query string, args ...interface{}) DbStorageScannableResultSingleRow {
	return t.tx.QueryRowx(query, args...)
}
func (t *txWrap) Exec(query string, args ...interface{}) (DbStorageExecResult, error) {
	return t.tx.Exec(query, args...)
}
func (t *txWrap) MustExec(query string, args ...interface{}) DbStorageExecResult {
	return t.tx.MustExec(query, args...)
}
func (t *txWrap) Commit() error {
	return t.tx.Commit()
}
func (t *txWrap) Rollback() error {
	return t.tx.Rollback()
}
