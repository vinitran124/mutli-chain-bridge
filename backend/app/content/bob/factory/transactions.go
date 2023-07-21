// Code generated by BobGen psql v0.21.1. DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package factory

import (
	"context"
	"time"

	models "bridge/app/content/bob"
	"github.com/aarondl/opt/omit"
	"github.com/google/uuid"
	"github.com/jaswdr/faker"
	"github.com/stephenafamo/bob"
)

type TransactionMod interface {
	Apply(*TransactionTemplate)
}

type TransactionModFunc func(*TransactionTemplate)

func (f TransactionModFunc) Apply(n *TransactionTemplate) {
	f(n)
}

type TransactionModSlice []TransactionMod

func (mods TransactionModSlice) Apply(n *TransactionTemplate) {
	for _, f := range mods {
		f.Apply(n)
	}
}

// TransactionTemplate is an object representing the database table.
// all columns are optional and should be set by mods
type TransactionTemplate struct {
	ID           func() uuid.UUID
	User         func() string
	Token        func() string
	RawAmount    func() string
	ChainID      func() string
	IsWithdrawal func() bool
	CreatedAt    func() time.Time
	UpdatedAt    func() time.Time
	Hash         func() string

	f *factory
}

// Apply mods to the TransactionTemplate
func (o *TransactionTemplate) Apply(mods ...TransactionMod) {
	for _, mod := range mods {
		mod.Apply(o)
	}
}

// toModel returns an *models.Transaction
// this does nothing with the relationship templates
func (o TransactionTemplate) toModel() *models.Transaction {
	m := &models.Transaction{}

	if o.ID != nil {
		m.ID = o.ID()
	}
	if o.User != nil {
		m.User = o.User()
	}
	if o.Token != nil {
		m.Token = o.Token()
	}
	if o.RawAmount != nil {
		m.RawAmount = o.RawAmount()
	}
	if o.ChainID != nil {
		m.ChainID = o.ChainID()
	}
	if o.IsWithdrawal != nil {
		m.IsWithdrawal = o.IsWithdrawal()
	}
	if o.CreatedAt != nil {
		m.CreatedAt = o.CreatedAt()
	}
	if o.UpdatedAt != nil {
		m.UpdatedAt = o.UpdatedAt()
	}
	if o.Hash != nil {
		m.Hash = o.Hash()
	}

	return m
}

// toModels returns an models.TransactionSlice
// this does nothing with the relationship templates
func (o TransactionTemplate) toModels(number int) models.TransactionSlice {
	m := make(models.TransactionSlice, number)

	for i := range m {
		m[i] = o.toModel()
	}

	return m
}

// setModelRels creates and sets the relationships on *models.Transaction
// according to the relationships in the template. Nothing is inserted into the db
func (t TransactionTemplate) setModelRels(o *models.Transaction) {}

// BuildSetter returns an *models.TransactionSetter
// this does nothing with the relationship templates
func (o TransactionTemplate) BuildSetter() *models.TransactionSetter {
	m := &models.TransactionSetter{}

	if o.ID != nil {
		m.ID = omit.From(o.ID())
	}
	if o.User != nil {
		m.User = omit.From(o.User())
	}
	if o.Token != nil {
		m.Token = omit.From(o.Token())
	}
	if o.RawAmount != nil {
		m.RawAmount = omit.From(o.RawAmount())
	}
	if o.ChainID != nil {
		m.ChainID = omit.From(o.ChainID())
	}
	if o.IsWithdrawal != nil {
		m.IsWithdrawal = omit.From(o.IsWithdrawal())
	}
	if o.CreatedAt != nil {
		m.CreatedAt = omit.From(o.CreatedAt())
	}
	if o.UpdatedAt != nil {
		m.UpdatedAt = omit.From(o.UpdatedAt())
	}
	if o.Hash != nil {
		m.Hash = omit.From(o.Hash())
	}

	return m
}

// BuildManySetter returns an []*models.TransactionSetter
// this does nothing with the relationship templates
func (o TransactionTemplate) BuildManySetter(number int) []*models.TransactionSetter {
	m := make([]*models.TransactionSetter, number)

	for i := range m {
		m[i] = o.BuildSetter()
	}

	return m
}

// Build returns an *models.Transaction
// Related objects are also created and placed in the .R field
// NOTE: Objects are not inserted into the database. Use TransactionTemplate.Create
func (o TransactionTemplate) Build() *models.Transaction {
	m := o.toModel()
	o.setModelRels(m)

	return m
}

// BuildMany returns an models.TransactionSlice
// Related objects are also created and placed in the .R field
// NOTE: Objects are not inserted into the database. Use TransactionTemplate.CreateMany
func (o TransactionTemplate) BuildMany(number int) models.TransactionSlice {
	m := make(models.TransactionSlice, number)

	for i := range m {
		m[i] = o.Build()
	}

	return m
}

func ensureCreatableTransaction(m *models.TransactionSetter) {
	if m.ID.IsUnset() {
		m.ID = omit.From(random[uuid.UUID](nil))
	}
	if m.User.IsUnset() {
		m.User = omit.From(random[string](nil))
	}
	if m.Token.IsUnset() {
		m.Token = omit.From(random[string](nil))
	}
	if m.RawAmount.IsUnset() {
		m.RawAmount = omit.From(random[string](nil))
	}
	if m.ChainID.IsUnset() {
		m.ChainID = omit.From(random[string](nil))
	}
	if m.CreatedAt.IsUnset() {
		m.CreatedAt = omit.From(random[time.Time](nil))
	}
	if m.UpdatedAt.IsUnset() {
		m.UpdatedAt = omit.From(random[time.Time](nil))
	}
	if m.Hash.IsUnset() {
		m.Hash = omit.From(random[string](nil))
	}
}

// insertOptRels creates and inserts any optional the relationships on *models.Transaction
// according to the relationships in the template.
// any required relationship should have already exist on the model
func (o *TransactionTemplate) insertOptRels(ctx context.Context, exec bob.Executor, m *models.Transaction) (context.Context, error) {
	var err error

	return ctx, err
}

// Create builds a transaction and inserts it into the database
// Relations objects are also inserted and placed in the .R field
func (o *TransactionTemplate) Create(ctx context.Context, exec bob.Executor) (*models.Transaction, error) {
	_, m, err := o.create(ctx, exec)
	return m, err
}

// create builds a transaction and inserts it into the database
// Relations objects are also inserted and placed in the .R field
// this returns a context that includes the newly inserted model
func (o *TransactionTemplate) create(ctx context.Context, exec bob.Executor) (context.Context, *models.Transaction, error) {
	var err error
	opt := o.BuildSetter()
	ensureCreatableTransaction(opt)

	m, err := models.TransactionsTable.Insert(ctx, exec, opt)
	if err != nil {
		return ctx, nil, err
	}
	ctx = transactionCtx.WithValue(ctx, m)

	ctx, err = o.insertOptRels(ctx, exec, m)
	return ctx, m, err
}

// CreateMany builds multiple transactions and inserts them into the database
// Relations objects are also inserted and placed in the .R field
func (o TransactionTemplate) CreateMany(ctx context.Context, exec bob.Executor, number int) (models.TransactionSlice, error) {
	_, m, err := o.createMany(ctx, exec, number)
	return m, err
}

// createMany builds multiple transactions and inserts them into the database
// Relations objects are also inserted and placed in the .R field
// this returns a context that includes the newly inserted models
func (o TransactionTemplate) createMany(ctx context.Context, exec bob.Executor, number int) (context.Context, models.TransactionSlice, error) {
	var err error
	m := make(models.TransactionSlice, number)

	for i := range m {
		ctx, m[i], err = o.create(ctx, exec)
		if err != nil {
			return ctx, nil, err
		}
	}

	return ctx, m, nil
}

// Transaction has methods that act as mods for the TransactionTemplate
var TransactionMods transactionMods

type transactionMods struct{}

func (m transactionMods) RandomizeAllColumns(f *faker.Faker) TransactionMod {
	return TransactionModSlice{
		TransactionMods.RandomID(f),
		TransactionMods.RandomUser(f),
		TransactionMods.RandomToken(f),
		TransactionMods.RandomRawAmount(f),
		TransactionMods.RandomChainID(f),
		TransactionMods.RandomIsWithdrawal(f),
		TransactionMods.RandomCreatedAt(f),
		TransactionMods.RandomUpdatedAt(f),
		TransactionMods.RandomHash(f),
	}
}

// Set the model columns to this value
func (m transactionMods) ID(val uuid.UUID) TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		o.ID = func() uuid.UUID { return val }
	})
}

// Set the Column from the function
func (m transactionMods) IDFunc(f func() uuid.UUID) TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		o.ID = f
	})
}

// Clear any values for the column
func (m transactionMods) UnsetID() TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		o.ID = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m transactionMods) RandomID(f *faker.Faker) TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		o.ID = func() uuid.UUID {
			return random[uuid.UUID](f)
		}
	})
}

func (m transactionMods) ensureID(f *faker.Faker) TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		if o.ID != nil {
			return
		}

		o.ID = func() uuid.UUID {
			return random[uuid.UUID](f)
		}
	})
}

// Set the model columns to this value
func (m transactionMods) User(val string) TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		o.User = func() string { return val }
	})
}

// Set the Column from the function
func (m transactionMods) UserFunc(f func() string) TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		o.User = f
	})
}

// Clear any values for the column
func (m transactionMods) UnsetUser() TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		o.User = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m transactionMods) RandomUser(f *faker.Faker) TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		o.User = func() string {
			return random[string](f)
		}
	})
}

func (m transactionMods) ensureUser(f *faker.Faker) TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		if o.User != nil {
			return
		}

		o.User = func() string {
			return random[string](f)
		}
	})
}

// Set the model columns to this value
func (m transactionMods) Token(val string) TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		o.Token = func() string { return val }
	})
}

// Set the Column from the function
func (m transactionMods) TokenFunc(f func() string) TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		o.Token = f
	})
}

// Clear any values for the column
func (m transactionMods) UnsetToken() TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		o.Token = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m transactionMods) RandomToken(f *faker.Faker) TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		o.Token = func() string {
			return random[string](f)
		}
	})
}

func (m transactionMods) ensureToken(f *faker.Faker) TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		if o.Token != nil {
			return
		}

		o.Token = func() string {
			return random[string](f)
		}
	})
}

// Set the model columns to this value
func (m transactionMods) RawAmount(val string) TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		o.RawAmount = func() string { return val }
	})
}

// Set the Column from the function
func (m transactionMods) RawAmountFunc(f func() string) TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		o.RawAmount = f
	})
}

// Clear any values for the column
func (m transactionMods) UnsetRawAmount() TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		o.RawAmount = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m transactionMods) RandomRawAmount(f *faker.Faker) TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		o.RawAmount = func() string {
			return random[string](f)
		}
	})
}

func (m transactionMods) ensureRawAmount(f *faker.Faker) TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		if o.RawAmount != nil {
			return
		}

		o.RawAmount = func() string {
			return random[string](f)
		}
	})
}

// Set the model columns to this value
func (m transactionMods) ChainID(val string) TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		o.ChainID = func() string { return val }
	})
}

// Set the Column from the function
func (m transactionMods) ChainIDFunc(f func() string) TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		o.ChainID = f
	})
}

// Clear any values for the column
func (m transactionMods) UnsetChainID() TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		o.ChainID = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m transactionMods) RandomChainID(f *faker.Faker) TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		o.ChainID = func() string {
			return random[string](f)
		}
	})
}

func (m transactionMods) ensureChainID(f *faker.Faker) TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		if o.ChainID != nil {
			return
		}

		o.ChainID = func() string {
			return random[string](f)
		}
	})
}

// Set the model columns to this value
func (m transactionMods) IsWithdrawal(val bool) TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		o.IsWithdrawal = func() bool { return val }
	})
}

// Set the Column from the function
func (m transactionMods) IsWithdrawalFunc(f func() bool) TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		o.IsWithdrawal = f
	})
}

// Clear any values for the column
func (m transactionMods) UnsetIsWithdrawal() TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		o.IsWithdrawal = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m transactionMods) RandomIsWithdrawal(f *faker.Faker) TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		o.IsWithdrawal = func() bool {
			return random[bool](f)
		}
	})
}

func (m transactionMods) ensureIsWithdrawal(f *faker.Faker) TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		if o.IsWithdrawal != nil {
			return
		}

		o.IsWithdrawal = func() bool {
			return random[bool](f)
		}
	})
}

// Set the model columns to this value
func (m transactionMods) CreatedAt(val time.Time) TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		o.CreatedAt = func() time.Time { return val }
	})
}

// Set the Column from the function
func (m transactionMods) CreatedAtFunc(f func() time.Time) TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		o.CreatedAt = f
	})
}

// Clear any values for the column
func (m transactionMods) UnsetCreatedAt() TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		o.CreatedAt = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m transactionMods) RandomCreatedAt(f *faker.Faker) TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		o.CreatedAt = func() time.Time {
			return random[time.Time](f)
		}
	})
}

func (m transactionMods) ensureCreatedAt(f *faker.Faker) TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		if o.CreatedAt != nil {
			return
		}

		o.CreatedAt = func() time.Time {
			return random[time.Time](f)
		}
	})
}

// Set the model columns to this value
func (m transactionMods) UpdatedAt(val time.Time) TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		o.UpdatedAt = func() time.Time { return val }
	})
}

// Set the Column from the function
func (m transactionMods) UpdatedAtFunc(f func() time.Time) TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		o.UpdatedAt = f
	})
}

// Clear any values for the column
func (m transactionMods) UnsetUpdatedAt() TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		o.UpdatedAt = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m transactionMods) RandomUpdatedAt(f *faker.Faker) TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		o.UpdatedAt = func() time.Time {
			return random[time.Time](f)
		}
	})
}

func (m transactionMods) ensureUpdatedAt(f *faker.Faker) TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		if o.UpdatedAt != nil {
			return
		}

		o.UpdatedAt = func() time.Time {
			return random[time.Time](f)
		}
	})
}

// Set the model columns to this value
func (m transactionMods) Hash(val string) TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		o.Hash = func() string { return val }
	})
}

// Set the Column from the function
func (m transactionMods) HashFunc(f func() string) TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		o.Hash = f
	})
}

// Clear any values for the column
func (m transactionMods) UnsetHash() TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		o.Hash = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m transactionMods) RandomHash(f *faker.Faker) TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		o.Hash = func() string {
			return random[string](f)
		}
	})
}

func (m transactionMods) ensureHash(f *faker.Faker) TransactionMod {
	return TransactionModFunc(func(o *TransactionTemplate) {
		if o.Hash != nil {
			return
		}

		o.Hash = func() string {
			return random[string](f)
		}
	})
}
