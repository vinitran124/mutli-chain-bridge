package etherman

// Config represents the configuration of the etherman
type Config struct {
	BscTestnet ChainConfig `mapstructure:"BscTestnet"`
	Sepolia    ChainConfig `mapstructure:"Sepolia"`
}

type ChainConfig struct {
	RPC            string `mapstructure:"RPC"`
	ChainId        uint64 `mapstructure:"ChainId"`
	Erc20TokenList string `mapstructure:"Erc20TokenList"`
	BridgeAddress  string `mapstructure:"BridgeAddress"`
	PrivateKey     string `mapstructure:"PrivateKey"`
}

const (
	BSC_TESTNET_CHAIN_ID = 97
	SEPOLIA_CHAIN_ID     = 11155111
)
