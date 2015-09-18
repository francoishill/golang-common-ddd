package DatabaseSqlxDbStorage

import (
	"database/sql"
	. "github.com/francoishill/golang-web-dry/errors/checkerror"
	"github.com/jmoiron/sqlx"
	"github.com/rubenv/sql-migrate"

	. "github.com/francoishill/golang-common-ddd/Interface/Storage/DbStorage"
)

type storage struct {
	db            *sqlx.DB
	migrationPath string
}

func (s *storage) Migrate() (numMigrationsApplied int) {
	migrations := &migrate.FileMigrationSource{
		Dir: s.migrationPath,
	}

	numMigrationsApplied, err := migrate.Exec(s.db.DB, s.db.DriverName(), migrations, migrate.Up)
	CheckError(err)
	return numMigrationsApplied
}

func (s *storage) BeginTransaction() (DbStorageTransaction, error) {
	return s.db.Beginx()
}
func (s *storage) MustBeginTransaction() DbStorageTransaction {
	return s.db.MustBegin()
}

func (s *storage) Select(dest interface{}, query string, args ...interface{}) error {
	return s.db.Select(dest, query, args...)
}
func (s *storage) Get(dest interface{}, query string, args ...interface{}) error {
	return s.db.Get(dest, query, args...)
}

func (s *storage) Query(query string, args ...interface{}) (DbStorageScannableResultMultipleRows, error) {
	return s.db.Queryx(query, args...)
}

func (s *storage) QueryRow(query string, args ...interface{}) DbStorageScannableResultSingleRow {
	return s.db.QueryRowx(query, args...)
}

func (s *storage) Exec(query string, args ...interface{}) (DbStorageExecResult, error) {
	return s.db.Exec(query, args...)
}
func (s *storage) MustExec(query string, args ...interface{}) DbStorageExecResult {
	return s.db.MustExec(query, args...)
}

func (s *storage) DeferableCommitOnSuccessRollbackOnFail(tx DbStorageTransaction) {
	if r := recover(); r != nil {
		tx.Rollback()
		panic(r)
	} else {
		err := tx.Commit()
		CheckError(err)
	}
}

func (s *storage) IsNoRowsError(err error) bool {
	return err == sql.ErrNoRows
}

func New(driverName, dataSourceName, migrationPath string) DbStorage {
	return &storage{
		sqlx.MustConnect(driverName, dataSourceName),
		migrationPath,
	}
}
