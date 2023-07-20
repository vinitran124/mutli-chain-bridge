package blockchain

import (
	"bridge/app/blockchain/chainFactory"
	"bridge/app/blockchain/ethereum"
)

func NewChain(chainId string) (chainFactory.IChain, error) {
	chain := new(ethereum.Ethereum)
	err := chain.Init(chainId)
	if err != nil {
		return nil, err
	}

	err = chain.NewClient()
	if err != nil {
		return nil, err
	}

	return chain, nil
}
