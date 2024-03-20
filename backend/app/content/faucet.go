package content

import (
	"bridge/app/content/bob"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

const amountFaucet = "10000000000000000000" //10token

type FaucetRequest struct {
	UserAddress string `json:"user_address"`
	Token       string `json:"token"`
	ChainId     string `json:"chain_id"`
}

func (v *V1Router) faucet(c *gin.Context) {
	ctx := context.Background()
	var request FaucetRequest
	err := c.BindJSON(&request)
	if err != nil {
		responseErrUnauthorized(c)
		return
	}

	request.Token = strings.ToLower(request.Token)
	request.UserAddress = strings.ToLower(request.UserAddress)

	isValidToken, err := bob.Tokens(
		ctx,
		SQLRepository(),
		bob.SelectWhere.Tokens.Address.EQ(request.Token),
		bob.SelectWhere.Tokens.ChainID.EQ(request.ChainId)).Exists()
	if err != nil {
		responseFailureWithMessage(c, "invalid token")
		return
	}

	if isValidToken == false {
		responseFailureWithMessage(c, "invalid token")
		return
	}

	txHash, err := ChainRepository(request.ChainId).TransferErc20(request.Token, request.UserAddress, amountFaucet)
	if err != nil {
		responseFailureWithMessage(c, fmt.Sprintln(err))
	}

	responseSuccessWithMessage(c, fmt.Sprintf("Tx Hash: %s", txHash))
}
