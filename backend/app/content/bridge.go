package content

import (
	"bridge/app/content/bob"
	"bridge/app/content/datastore"
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"math/big"
	"strings"
)

type BridgeRequest struct {
	InChain      string `json:"in_chain"`
	OutChain     string `json:"out_chain"`
	Amount       string `json:"amount"`
	TokenAddress string `json:"token_address"`
	UserAddress  string `json:"user_address"`
}

func (v *V1Router) bridge(c *gin.Context) {
	ctx := context.Background()
	var auth BridgeRequest
	err := c.BindJSON(&auth)
	if err != nil {
		responseErrUnauthorized(c)
		return
	}
	auth.TokenAddress = strings.ToLower(auth.TokenAddress)
	auth.UserAddress = strings.ToLower(auth.UserAddress)

	if auth.OutChain == auth.InChain {
		responseFailureWithMessage(c, "invalid input and output chainId")
		return
	}

	tokenIn, err := bob.Tokens(
		ctx,
		SQLRepository(),
		bob.SelectWhere.Tokens.Address.EQ(auth.TokenAddress),
		bob.SelectWhere.Tokens.ChainID.EQ(auth.InChain)).One()
	if err != nil {
		responseFailureWithMessage(c, "invalid input token")
		return
	}

	tokenOut, err := bob.Tokens(
		ctx,
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

	dataStr := datastore.DatastoreBridgeRequest{}

	inRequest, err := dataStr.IsInRequest(ctx, SQLRepository(), auth.UserAddress)
	if inRequest == true {
		responseFailureWithMessage(c, "you have transaction in progress, please waiting")
		return
	}

	tx, err := SQLRepository().BeginTx(ctx, &sql.TxOptions{})
	bridgeRq, err := dataStr.Insert(ctx, tx, &bob.BridgeRequest{
		InputChain:  auth.InChain,
		OutputChain: auth.OutChain,
		RawAmount:   auth.Amount,
		Token:       auth.TokenAddress,
		UserAddress: auth.UserAddress,
	})
	if err != nil {
		log.Println(err)
		responseErrInternalServerError(c)
		return
	}

	err = tx.Commit()
	if err != nil {
		responseErrInternalServerError(c)
		return
	}

	responseSuccess(c, bridgeRq.ID.String())
}
