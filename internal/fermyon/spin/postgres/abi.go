// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

package postgres

import (
	rdbmstypes "github.com/fermyon/spin-go-sdk/internal/fermyon/spin/rdbms-types"
	"unsafe"
)

// RowSetShape is used for storage in variant or result types.
type RowSetShape struct {
	shape [unsafe.Sizeof(rdbmstypes.RowSet{})]byte
}

// PgErrorShape is used for storage in variant or result types.
type PgErrorShape struct {
	shape [unsafe.Sizeof(PgError{})]byte
}