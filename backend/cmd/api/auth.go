package api

import (
	"github.com/gin-gonic/gin"
)

type AuthRequest struct {
	InChain      string `json:"in_chain"`
	OutChain     string `json:"out_chain"`
	Amount       string `json:"amount"`
	TokenAddress string `json:"token_address"`
}

func (v *V1Router) auth(c *gin.Context) {
	var auth AuthRequest
	err := c.BindJSON(&auth)
	if err != nil {
		responseErrUnauthorized(c)
		return
	}

	//inChain := ChainRepository("5555")

	//a := inChain.GetTokenName("0x596c14ba2b6e759c73895ce13e64e49054181a7f")
	responseSuccessWithMessage(c, "a")
}
