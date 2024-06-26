// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	DeleteAccount(ctx context.Context, id int64) error
	GetAccount(ctx context.Context, id int64) (Account, error)
	ListAccoutns(ctx context.Context, arg ListAccoutnsParams) ([]Account, error)
	UppdateAccount(ctx context.Context, arg UppdateAccountParams) (Account, error)
}

var _ Querier = (*Queries)(nil)
