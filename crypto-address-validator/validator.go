package crypto_address_validator

import (
	"strings"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/cpacia/bchutil"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

//NetworkType ネットワーク種別
type NetworkType string

// network_type
const (
	NetworkTypeMainNet  NetworkType = "mainnet"
	NetworkTypeTestNet3 NetworkType = "testnet3"
	//NetworkTypeRegTestNet NetworkType = "regtest"
)

type CoinType string

const (
	CoinTypeBTC CoinType = "BTC"
	CoinTypeBCH CoinType = "BCH"
)

var (
	BTCPrefix = map[CoinType]map[NetworkType][]string{
		CoinTypeBTC: {
			NetworkTypeMainNet:  {"1", "3", "xpub", "bc1"},
			NetworkTypeTestNet3: {"m", "n", "2", "tpub", "tb1"},
		},
		CoinTypeBCH: {
			NetworkTypeMainNet:  {"bitcoincash", "1", "3", "xpub"},
			NetworkTypeTestNet3: {"bchtest", "m", "n", "2", "tpub"},
		},
	}
)

// getChain returns chain information for bitcoin, bitcoin cash
func getChain(netType NetworkType) *chaincfg.Params {
	if netType == NetworkTypeMainNet {
		return &chaincfg.MainNetParams
	}

	return &chaincfg.TestNet3Params
}

func checkPrefix(addr string, netType NetworkType, coinType CoinType) bool {
	for _, val := range BTCPrefix[coinType][netType] {
		if strings.Index(addr, val) == 0 {
			return true
		}
	}
	return false
}

// ValidateBTCAddr is to validate BTC address and expected to be called from cayenne platform
func ValidateBTCAddr(addr string, netType NetworkType) error {
	// decode
	_, err := DecodeBTCAddr(addr, getChain(netType))
	if err != nil {
		return errors.Errorf("address [%s] is invalid", addr)
	}

	//mainnet/testnetのバージョンの違い
	if !checkPrefix(addr, netType, CoinTypeBTC) {
		return errors.Errorf("address [%s] is invalid because of difference of network", addr)
	}

	return nil
}

// decodeBTCAddr is to decode BTC address from string
func DecodeBTCAddr(addr string, chainCfg *chaincfg.Params) (btcutil.Address, error) {
	address, err := btcutil.DecodeAddress(addr, chainCfg)
	if err != nil {
		return nil, errors.Errorf("address [%s] is invalid", addr)
	}

	return address, nil
}

// ValidateBCHAddr is to validate BCH address and expected to be called from cayenne platform
func ValidateBCHAddr(addr string, netType NetworkType) error {
	// decode
	_, err := DecodeBCHAddr(addr, getChain(netType))
	if err != nil {
		return errors.Errorf("address [%s] is invalid", addr)
	}

	//mainnet/testnetのバージョンの違い
	if !checkPrefix(addr, netType, CoinTypeBCH) {
		return errors.Errorf("address [%s] is invalid because of difference of network", addr)
	}

	return nil
}

// decodeBTCAddr is to decode BTC address from string
func DecodeBCHAddr(addr string, chainCfg *chaincfg.Params) (btcutil.Address, error) {
	//Legacyアドレスはerrorを返す
	address, err := bchutil.DecodeAddress(addr, chainCfg)
	if err != nil {
		//Note: when new cashAddr is given, that above DecodeAddress() will fail
		//FIXME: testnet接続時でもmainnetのLegacyアドレスのValiationが通ってしまう
		address, err = btcutil.DecodeAddress(addr, chainCfg)
		if err != nil {
			return nil, errors.Errorf("address [%s] is invalid", addr)
		}
	}

	return address, nil
}

// ValidateETHAddr is to validate ETH address and expected to be called from cayenne platform
func ValidateETHAddr(addr string) error {
	// validation check
	if !common.IsHexAddress(addr) {
		return errors.Errorf("address [%s] is invalid", addr)
	}
	return nil
}
