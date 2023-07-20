package main

import (
	"bridge/app/utils"
	"github.com/uptrace/bun"
)

func SQLRepository() *bun.DB {
	return utils.GetContextSQL().Client
}
