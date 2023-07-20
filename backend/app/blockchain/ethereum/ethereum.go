package ethereum

import (
	"bridge/app/blockchain/chainFactory"
	"bridge/app/model"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"os"
)

type Ethereum struct {
	chainFactory.Chain
	client       *ethclient.Client
	bridgeRouter *BridgeRouter
}

func (e *Ethereum) NewClient() error {
	chainRpc := os.Getenv(fmt.Sprintf("RPC_%s", e.ChainId))
	if chainRpc == "" {
		return fmt.Errorf("chain: Invalid chain id")
	}

	client, err := ethclient.Dial(chainRpc)
	if err != nil {
		return err
	}

	bridgeContract := common.HexToAddress(os.Getenv(fmt.Sprintf("%s%s", chainFactory.EnvBridgeContract, e.ChainId)))
	bridgeRouter, err := NewBridgeRouter(bridgeContract, client)
	if err != nil {
		return err
	}

	e.bridgeRouter = bridgeRouter
	e.client = client
	return nil
}

func (e *Ethereum) TrackDeposit(events chan model.DepositEvent) error {
	bridgeContract := common.HexToAddress(os.Getenv(fmt.Sprintf("%s%s", chainFactory.EnvBridgeContract, e.ChainId)))
	e.getEvents(events, bridgeContract)
	return nil
}

func (e *Ethereum) CallWithdrawal(token, user, amount string) error {
	privateKey := os.Getenv(fmt.Sprintf("%s%s", chainFactory.EnvPrivateKey, e.ChainId))
	auth, err := e.Auth(privateKey)
	if err != nil {
		return err
	}

	err = e.CallWithdrawalInContract(auth, token, user, StringToBigInt(amount))
	return nil
}

func (e *Ethereum) GetTokenInPool(token string) (*big.Int, error) {
	return e.GetTokenAmountInPool(token)
}
