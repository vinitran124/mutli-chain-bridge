// Code generated by BobGen psql v0.21.1. DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package bob

import (
	"context"
	"time"

	"github.com/aarondl/opt/omit"
	"github.com/google/uuid"
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/dialect"
	"github.com/stephenafamo/bob/dialect/psql/sm"
)

// BridgeRequest is an object representing the database table.
type BridgeRequest struct {
	ID          uuid.UUID `db:"id,pk" `
	UserAddress string    `db:"user_address" `
	InputChain  string    `db:"input_chain" `
	OutputChain string    `db:"output_chain" `
	RawAmount   string    `db:"raw_amount" `
	IsComplete  bool      `db:"is_complete" `
	Token       string    `db:"token" `
	CreatedAt   time.Time `db:"created_at" `
	UpdatedAt   time.Time `db:"updated_at" `
}

// BridgeRequestSlice is an alias for a slice of pointers to BridgeRequest.
// This should almost always be used instead of []*BridgeRequest.
type BridgeRequestSlice []*BridgeRequest

// BridgeRequestsTable contains methods to work with the bridge_requests table
var BridgeRequestsTable = psql.NewTablex[*BridgeRequest, BridgeRequestSlice, *BridgeRequestSetter]("", "bridge_requests")

// BridgeRequestsQuery is a query on the bridge_requests table
type BridgeRequestsQuery = *psql.TableQuery[*BridgeRequest, BridgeRequestSlice, *BridgeRequestSetter]

// BridgeRequestsStmt is a prepared statment on bridge_requests
type BridgeRequestsStmt = bob.QueryStmt[*BridgeRequest, BridgeRequestSlice]

// BridgeRequestSetter is used for insert/upsert/update operations
// All values are optional, and do not have to be set
// Generated columns are not included
type BridgeRequestSetter struct {
	ID          omit.Val[uuid.UUID] `db:"id,pk"`
	UserAddress omit.Val[string]    `db:"user_address"`
	InputChain  omit.Val[string]    `db:"input_chain"`
	OutputChain omit.Val[string]    `db:"output_chain"`
	RawAmount   omit.Val[string]    `db:"raw_amount"`
	IsComplete  omit.Val[bool]      `db:"is_complete"`
	Token       omit.Val[string]    `db:"token"`
	CreatedAt   omit.Val[time.Time] `db:"created_at"`
	UpdatedAt   omit.Val[time.Time] `db:"updated_at"`
}

type bridgeRequestColumnNames struct {
	ID          string
	UserAddress string
	InputChain  string
	OutputChain string
	RawAmount   string
	IsComplete  string
	Token       string
	CreatedAt   string
	UpdatedAt   string
}

var BridgeRequestColumns = struct {
	ID          psql.Expression
	UserAddress psql.Expression
	InputChain  psql.Expression
	OutputChain psql.Expression
	RawAmount   psql.Expression
	IsComplete  psql.Expression
	Token       psql.Expression
	CreatedAt   psql.Expression
	UpdatedAt   psql.Expression
}{
	ID:          psql.Quote("bridge_requests", "id"),
	UserAddress: psql.Quote("bridge_requests", "user_address"),
	InputChain:  psql.Quote("bridge_requests", "input_chain"),
	OutputChain: psql.Quote("bridge_requests", "output_chain"),
	RawAmount:   psql.Quote("bridge_requests", "raw_amount"),
	IsComplete:  psql.Quote("bridge_requests", "is_complete"),
	Token:       psql.Quote("bridge_requests", "token"),
	CreatedAt:   psql.Quote("bridge_requests", "created_at"),
	UpdatedAt:   psql.Quote("bridge_requests", "updated_at"),
}

type bridgeRequestWhere[Q psql.Filterable] struct {
	ID          psql.WhereMod[Q, uuid.UUID]
	UserAddress psql.WhereMod[Q, string]
	InputChain  psql.WhereMod[Q, string]
	OutputChain psql.WhereMod[Q, string]
	RawAmount   psql.WhereMod[Q, string]
	IsComplete  psql.WhereMod[Q, bool]
	Token       psql.WhereMod[Q, string]
	CreatedAt   psql.WhereMod[Q, time.Time]
	UpdatedAt   psql.WhereMod[Q, time.Time]
}

func BridgeRequestWhere[Q psql.Filterable]() bridgeRequestWhere[Q] {
	return bridgeRequestWhere[Q]{
		ID:          psql.Where[Q, uuid.UUID](BridgeRequestColumns.ID),
		UserAddress: psql.Where[Q, string](BridgeRequestColumns.UserAddress),
		InputChain:  psql.Where[Q, string](BridgeRequestColumns.InputChain),
		OutputChain: psql.Where[Q, string](BridgeRequestColumns.OutputChain),
		RawAmount:   psql.Where[Q, string](BridgeRequestColumns.RawAmount),
		IsComplete:  psql.Where[Q, bool](BridgeRequestColumns.IsComplete),
		Token:       psql.Where[Q, string](BridgeRequestColumns.Token),
		CreatedAt:   psql.Where[Q, time.Time](BridgeRequestColumns.CreatedAt),
		UpdatedAt:   psql.Where[Q, time.Time](BridgeRequestColumns.UpdatedAt),
	}
}

// BridgeRequests begins a query on bridge_requests
func BridgeRequests(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) BridgeRequestsQuery {
	return BridgeRequestsTable.Query(ctx, exec, mods...)
}

// FindBridgeRequest retrieves a single record by primary key
// If cols is empty Find will return all columns.
func FindBridgeRequest(ctx context.Context, exec bob.Executor, IDPK uuid.UUID, cols ...string) (*BridgeRequest, error) {
	if len(cols) == 0 {
		return BridgeRequestsTable.Query(
			ctx, exec,
			SelectWhere.BridgeRequests.ID.EQ(IDPK),
		).One()
	}

	return BridgeRequestsTable.Query(
		ctx, exec,
		SelectWhere.BridgeRequests.ID.EQ(IDPK),
		sm.Columns(BridgeRequestsTable.Columns().Only(cols...)),
	).One()
}

// BridgeRequestExists checks the presence of a single record by primary key
func BridgeRequestExists(ctx context.Context, exec bob.Executor, IDPK uuid.UUID) (bool, error) {
	return BridgeRequestsTable.Query(
		ctx, exec,
		SelectWhere.BridgeRequests.ID.EQ(IDPK),
	).Exists()
}

// Update uses an executor to update the BridgeRequest
func (o *BridgeRequest) Update(ctx context.Context, exec bob.Executor, cols ...string) (int64, error) {
	rowsAff, err := BridgeRequestsTable.Update(ctx, exec, o, cols...)
	if err != nil {
		return rowsAff, err
	}

	return rowsAff, nil
}

// Delete deletes a single BridgeRequest record with an executor
func (o *BridgeRequest) Delete(ctx context.Context, exec bob.Executor) (int64, error) {
	return BridgeRequestsTable.Delete(ctx, exec, o)
}

// Reload refreshes the BridgeRequest using the executor
func (o *BridgeRequest) Reload(ctx context.Context, exec bob.Executor) error {
	o2, err := BridgeRequestsTable.Query(
		ctx, exec,
		SelectWhere.BridgeRequests.ID.EQ(o.ID),
	).One()
	if err != nil {
		return err
	}

	*o = *o2

	return nil
}

func (o BridgeRequestSlice) DeleteAll(ctx context.Context, exec bob.Executor) (int64, error) {
	return BridgeRequestsTable.DeleteMany(ctx, exec, o...)
}

func (o BridgeRequestSlice) UpdateAll(ctx context.Context, exec bob.Executor, vals BridgeRequestSetter) (int64, error) {
	rowsAff, err := BridgeRequestsTable.UpdateMany(ctx, exec, &vals, o...)
	if err != nil {
		return rowsAff, err
	}

	return rowsAff, nil
}

func (o BridgeRequestSlice) ReloadAll(ctx context.Context, exec bob.Executor) error {
	var mods []bob.Mod[*dialect.SelectQuery]

	IDPK := make([]uuid.UUID, len(o))

	for i, o := range o {
		IDPK[i] = o.ID
	}

	mods = append(mods,
		SelectWhere.BridgeRequests.ID.In(IDPK...),
	)

	o2, err := BridgeRequests(ctx, exec, mods...).All()
	if err != nil {
		return err
	}

	for _, old := range o {
		for _, new := range o2 {
			if new.ID != old.ID {
				continue
			}

			*old = *new
			break
		}
	}

	return nil
}