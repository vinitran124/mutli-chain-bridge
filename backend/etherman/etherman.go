package etherman

import (
	BridgeRouter "bridge/etherman/smartcontract/bridgeRouter"
	Erc20 "bridge/etherman/smartcontract/erc20token"
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
)

var (
	ErrNotFound           = errors.New("not found")
	ErrPrivateKeyNotFound = errors.New("can't find sender private key to sign tx")
)

type ethereumClient interface {
	ethereum.ChainReader
	ethereum.ChainStateReader
	ethereum.ContractCaller
	ethereum.GasEstimator
	ethereum.GasPricer
	ethereum.LogFilterer
	ethereum.TransactionReader
	ethereum.TransactionSender

	bind.DeployBackend
}

// Client is a simple implementation of EtherMan.
type Client struct {
	EthClient     ethereumClient
	Erc20Token    map[common.Address]*Erc20.Erc20Token
	BridgeRouter  *BridgeRouter.BridgeRouter
	cfg           ChainConfig
	sender        map[common.Address]*bind.TransactOpts // empty in case of read-only client
	SenderAddress []common.Address
}

func NewClientFromChainId(chainId uint64, cfg Config) (*Client, error) {
	switch chainId {
	case BSC_TESTNET_CHAIN_ID:
		return NewClient(cfg.BscTestnet)
	case SEPOLIA_CHAIN_ID:
		return NewClient(cfg.Sepolia)
	default:
		return nil, fmt.Errorf("invalid chain id")
	}
}

func NewClient(cfg ChainConfig) (*Client, error) {
	ethClient, err := ethclient.Dial(cfg.RPC)
	if err != nil {
		return nil, fmt.Errorf("error connecting to %s: %e", cfg.RPC, err)
	}

	erc20Token := make(map[common.Address]*Erc20.Erc20Token)
	erc20TokenList := strings.Split(cfg.Erc20TokenList, ".")
	for _, token := range erc20TokenList {
		tokenAddr := common.HexToAddress(token)
		instanceToken, err := Erc20.NewErc20Token(tokenAddr, ethClient)
		if err != nil {
			return nil, err
		}
		erc20Token[tokenAddr] = instanceToken
	}

	bridgeContract, err := BridgeRouter.NewBridgeRouter(common.HexToAddress(cfg.BridgeAddress), ethClient)
	if err != nil {
		return nil, err
	}

	sender := make(map[common.Address]*bind.TransactOpts)
	var senderAddress []common.Address

	privateKeyList := strings.Split(cfg.PrivateKey, ".")
	for _, privateKey := range privateKeyList {
		address, err := GetAddress(privateKey)
		if err != nil {
			return nil, err
		}
		senderAddress = append(senderAddress, address)

		auth, err := GetAuth(privateKey, cfg.ChainId)
		if err != nil {
			return nil, err
		}

		sender[address] = auth
	}
	return &Client{
		EthClient:     ethClient,
		Erc20Token:    erc20Token,
		cfg:           cfg,
		BridgeRouter:  bridgeContract,
		sender:        sender,
		SenderAddress: senderAddress,
	}, nil
}

// SendTx sends a tx
func (etherMan *Client) SendTx(ctx context.Context, tx *types.Transaction) error {
	return etherMan.EthClient.SendTransaction(ctx, tx)
}

// SignTx tries to sign a transaction accordingly to the provided sender
func (etherMan *Client) SignTx(sender common.Address, tx *types.Transaction) (*types.Transaction, error) {
	auth, err := etherMan.getAuthByAddress(sender)
	if err == ErrNotFound {
		return nil, ErrPrivateKeyNotFound
	}

	signedTx, err := auth.Signer(auth.From, tx)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

// getAuthByAddress tries to get an authorization from the authorizations map
func (etherMan *Client) getAuthByAddress(addr common.Address) (*bind.TransactOpts, error) {
	auth, found := etherMan.sender[addr]
	if !found {
		return nil, ErrNotFound
	}
	return auth, nil
}

// AddOrReplaceAuth adds an authorization or replace an existent one to the same account
func (etherMan *Client) AddOrReplaceAuth(auth bind.TransactOpts) error {
	etherMan.sender[auth.From] = &auth
	return nil
}

func (etherMan *Client) TransferErc20Token(account, token, to common.Address, amount *big.Int) (*types.Transaction, error) {
	opts, err := etherMan.getAuthByAddress(account)
	if err == ErrNotFound {
		return nil, errors.New("can't find account private key to sign tx")
	}

	tx, err := etherMan.Erc20Token[token].Transfer(opts, to, amount)
	if err != nil {
		return nil, fmt.Errorf("error transfer erc20 token to %s. Error: %w", to.String(), err)
	}

	return tx, nil
}
