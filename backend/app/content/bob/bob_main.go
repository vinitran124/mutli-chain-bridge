// Code generated by BobGen psql v0.21.1. DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package bob

import (
	"context"

	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/dialect"
)

var TableNames = struct {
	BridgeRequests  string
	GooseDBVersions string
	Tokens          string
	Transactions    string
}{
	BridgeRequests:  "bridge_requests",
	GooseDBVersions: "goose_db_version",
	Tokens:          "tokens",
	Transactions:    "transactions",
}

var ColumnNames = struct {
	BridgeRequests  bridgeRequestColumnNames
	GooseDBVersions gooseDBVersionColumnNames
	Tokens          tokenColumnNames
	Transactions    transactionColumnNames
}{
	BridgeRequests: bridgeRequestColumnNames{
		ID:          "id",
		UserAddress: "user_address",
		InputChain:  "input_chain",
		OutputChain: "output_chain",
		RawAmount:   "raw_amount",
		IsComplete:  "is_complete",
		Token:       "token",
		CreatedAt:   "created_at",
		UpdatedAt:   "updated_at",
	},
	GooseDBVersions: gooseDBVersionColumnNames{
		ID:        "id",
		VersionID: "version_id",
		IsApplied: "is_applied",
		Tstamp:    "tstamp",
	},
	Tokens: tokenColumnNames{
		ID:        "id",
		Address:   "address",
		ChainID:   "chain_id",
		Name:      "name",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
	},
	Transactions: transactionColumnNames{
		ID:         "id",
		User:       "user",
		Token:      "token",
		RawAmount:  "raw_amount",
		ChainID:    "chain_id",
		IsComplete: "is_complete",
		CreatedAt:  "created_at",
		UpdatedAt:  "updated_at",
		Hash:       "hash",
	},
}

var (
	SelectWhere = Where[*dialect.SelectQuery]()
	InsertWhere = Where[*dialect.InsertQuery]()
	UpdateWhere = Where[*dialect.UpdateQuery]()
	DeleteWhere = Where[*dialect.DeleteQuery]()
)

func Where[Q psql.Filterable]() struct {
	BridgeRequests  bridgeRequestWhere[Q]
	GooseDBVersions gooseDBVersionWhere[Q]
	Tokens          tokenWhere[Q]
	Transactions    transactionWhere[Q]
} {
	return struct {
		BridgeRequests  bridgeRequestWhere[Q]
		GooseDBVersions gooseDBVersionWhere[Q]
		Tokens          tokenWhere[Q]
		Transactions    transactionWhere[Q]
	}{
		BridgeRequests:  BridgeRequestWhere[Q](),
		GooseDBVersions: GooseDBVersionWhere[Q](),
		Tokens:          TokenWhere[Q](),
		Transactions:    TransactionWhere[Q](),
	}
}

var (
	SelectJoins = getJoins[*dialect.SelectQuery]
	UpdateJoins = getJoins[*dialect.UpdateQuery]
	DeleteJoins = getJoins[*dialect.DeleteQuery]
)

type joinSet[Q any] struct {
	InnerJoin Q
	LeftJoin  Q
	RightJoin Q
}

type joins[Q dialect.Joinable] struct {
}

func getJoins[Q dialect.Joinable](ctx context.Context) joins[Q] {
	return joins[Q]{}
}
