package datastore

import (
	b "bridge/app/content/bob"
	"context"
	"github.com/aarondl/opt/omit"
	"github.com/stephenafamo/bob"
)

type DatastoreTransaction struct{}

func (ds *DatastoreTransaction) InsertTransaction(ctx context.Context, exec bob.Executor, params *b.Transaction) (*b.Transaction, error) {
	paramsSetter := b.TransactionSetter{
		ID:           omit.From(params.ID),
		User:         omit.From(params.User),
		Token:        omit.From(params.Token),
		RawAmount:    omit.From(params.RawAmount),
		ChainID:      omit.From(params.ChainID),
		IsWithdrawal: omit.From(params.IsWithdrawal),
		Hash:         omit.From(params.Hash),
		CreatedAt:    omit.From(params.CreatedAt),
		UpdatedAt:    omit.From(params.UpdatedAt),
	}
	return b.TransactionsTable.Insert(ctx, exec, &paramsSetter)
}
