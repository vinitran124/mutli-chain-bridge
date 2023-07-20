package model

import (
	"github.com/uptrace/bun"
	"time"
)

type DepositEvent struct {
	bun.BaseModel `bun:"table:rank"`
	ID            int64     `bun:"id,pk,autoincrement" json:"id"`
	User          string    `bun:"user,notnull" json:"user" `
	Token         string    `bun:"token,notnull" json:"token" `
	RawAmount     string    `bun:"raw_amount,notnull" json:"rawAmount"`
	ChainId       string    `bun:"chain_id,notnull" json:"chainId"`
	IsWithdrawal  bool      `bun:"is_withdrawal,notnull" json:"isWithdrawal"`
	CreatedAt     time.Time `bun:"created_at,notnull" json:"created_at"`
}
