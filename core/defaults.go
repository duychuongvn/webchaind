package core

import (
	"math/big"
	"github.com/ethereumproject/go-ethereum/logger/glog"
	"github.com/ethereumproject/go-ethereum/core/assets"
	"encoding/json"
)

var (
	DefaultChainConfigID string
	DefaultTestnetChainConfigID string

	DefaultChainConfigName string
	DefaultTestnetChainConfigName string

	DefaultChainConfigChainID *big.Int
	DefaultTestnetChainConfigChainID *big.Int

	DefaultConfig *ChainConfig
	TestConfig *ChainConfig

	TestNetGenesis *GenesisDump
	DefaultGenesis *GenesisDump
)

func readConfigFromDefaults(configPath string) *SufficientChainConfig {
	data, err := assets.DEFAULTS.Open(configPath)
	if err != nil {
		glog.Fatalf("Err opening default chain config assets (%s): %v", configPath, err)
	}
	var config = &SufficientChainConfig{}
	if json.NewDecoder(data).Decode(config); err != nil {
		glog.Fatalf("%v", err)
	}
	return config
}

func init() {
	mainnetConfigDefaults := readConfigFromDefaults("/core/config/mainnet.json")
	mordenConfigDefaults := readConfigFromDefaults("/core/config/morden.json")

	DefaultChainConfigID = mainnetConfigDefaults.Identity
	DefaultTestnetChainConfigID = mordenConfigDefaults.Identity

	DefaultChainConfigName = mainnetConfigDefaults.Name
	DefaultTestnetChainConfigName = mordenConfigDefaults.Name

	DefaultChainConfigChainID = mainnetConfigDefaults.ChainConfig.GetChainID()
	DefaultTestnetChainConfigChainID = mordenConfigDefaults.ChainConfig.GetChainID()

	DefaultConfig = mainnetConfigDefaults.ChainConfig
	TestConfig = mordenConfigDefaults.ChainConfig

	DefaultGenesis = mainnetConfigDefaults.Genesis
	TestNetGenesis = mordenConfigDefaults.Genesis
}
