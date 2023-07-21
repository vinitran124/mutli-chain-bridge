package main

import (
	"bridge/app/content/bob"
	"context"
	"github.com/gin-gonic/gin"
	"math/big"
)

type BridgeRequest struct {
	InChain      string `json:"in_chain"`
	OutChain     string `json:"out_chain"`
	Amount       string `json:"amount"`
	TokenAddress string `json:"token_address"`
}

func (v *V1Router) auth(c *gin.Context) {
	var auth BridgeRequest
	err := c.BindJSON(&auth)
	if err != nil {
		responseErrUnauthorized(c)
		return
	}

	tokenIn, err := bob.Tokens(
		context.Background(),
		SQLRepository(),
		bob.SelectWhere.Tokens.Address.EQ(auth.TokenAddress),
		bob.SelectWhere.Tokens.ChainID.EQ(auth.InChain)).One()
	if err != nil {
		responseFailureWithMessage(c, "invalid input token")
		return
	}

	tokenOut, err := bob.Tokens(
		context.Background(),
		SQLRepository(),
		bob.SelectWhere.Tokens.Name.EQ(tokenIn.Name),
		bob.SelectWhere.Tokens.ChainID.EQ(auth.OutChain)).One()
	if err != nil {
		responseFailureWithMessage(c, "invalid output token")
		return
	}

	amountInPoolTokenOut, err := ChainRepository(auth.OutChain).GetTokenInPool(tokenOut.Address)
	if err != nil {
		responseErrInternalServerError(c)
		return
	}

	amountIn, ok := big.NewInt(0).SetString(auth.Amount, 10)
	if !ok {
		responseErrInternalServerError(c)
		return
	}

	// require amount output token in pool must be grater than amount input token
	if amountIn.Cmp(amountInPoolTokenOut) == 1 {
		responseFailureWithMessage(c, "amount output token is not enough")
		return
	}

	//a := inChain.GetTokenName("0x596c14ba2b6e759c73895ce13e64e49054181a7f")
	responseSuccessWithMessage(c, "a")
}
