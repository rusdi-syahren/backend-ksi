package shared

import (
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
)

// TransactionManager interface
type TransactionManager interface {
	Begin()
	Commit()
	Rollback()
	Manage(f func() error) (err error)
}

// TransactionManagerGorm struct
type TransactionManagerGorm struct {
	DBRead, DBWrite, Tx *gorm.DB
	sync.Mutex
}

// NewTransactionManagerGorm create new TransactionManagerGorm
func NewTransactionManagerGorm(dbRead, dbWrite *gorm.DB) *TransactionManagerGorm {
	return &TransactionManagerGorm{DBRead: dbRead, DBWrite: dbWrite}
}

// Begin will begin transaction on each child repository
func (r *TransactionManagerGorm) Begin() {
	r.Lock()
	r.Tx = r.DBWrite.Begin()
}

// Rollback will rollback data if error happened
func (r *TransactionManagerGorm) Rollback() {
	if r.Tx != nil {
		r.Tx.Rollback()
	}
	r.Tx = nil
	r.Unlock()
}

// Commit will commit transaction
func (r *TransactionManagerGorm) Commit() {
	if r.Tx != nil {
		r.Tx.Commit()
	}
	r.Tx = nil
	r.Unlock()
}

// Manage function will manage transaction
func (r *TransactionManagerGorm) Manage(f func() error) (err error) {
	r.Begin()
	defer func() {
		if rec := recover(); rec != nil {
			r.Rollback()
			err = fmt.Errorf("panic: %v", rec)
		} else if err != nil {
			r.Rollback()
		} else {
			r.Commit()
		}
	}()

	err = f()
	return
}
