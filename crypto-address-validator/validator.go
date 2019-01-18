package crypto_address_validator

import (
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

// DecodeBTCAddr is to decode BTC address from string
func DecodeBCHAddr(addr string, chainCfg *chaincfg.Params) (btcutil.Address, error) {
	address, err := bchutil.DecodeAddress(addr, chainCfg)
	if err != nil {
		address, err = btcutil.DecodeAddress(addr, chainCfg)
		if err != nil {
			return nil, errors.Errorf("address [%s] is invalid", addr)
		}
	}

	return address, nil
}

// DecodeBTCAddr is to decode BTC address from string
func DecodeBTCAddr(addr string, chainCfg *chaincfg.Params) (btcutil.Address, error) {
	address, err := btcutil.DecodeAddress(addr, chainCfg)
	if err != nil {
		return nil, errors.Errorf("address [%s] is invalid", addr)
	}

	return address, nil
}

// getChain returns chain information for bitcoin, bitcoin cash
func getChain(netType NetworkType) *chaincfg.Params {
	if netType == NetworkTypeMainNet {
		return &chaincfg.MainNetParams
	}

	return &chaincfg.TestNet3Params
}

// ValidateBTCAddr is to validate BTC address and expected to be called from cayenne platform
func ValidateBTCAddr(addr string, netType NetworkType) error {
	// decode
	_, err := btcutil.DecodeAddress(addr, getChain(netType))
	if err != nil {
		return errors.Errorf("address [%s] is invalid", addr)
	}
	return nil
}

// ValidateBCHAddr is to validate BCH address and expected to be called from cayenne platform
func ValidateBCHAddr(addr string, netType NetworkType) error {
	// decode
	_, err := DecodeBCHAddr(addr, getChain(netType))
	if err != nil {
		return errors.Errorf("address [%s] is invalid", addr)
	}
	return nil
}

// ValidateETHAddr is to validate ETH address and expected to be called from cayenne platform
func ValidateETHAddr(addr string) error {
	// validation check
	if !common.IsHexAddress(addr) {
		return errors.Errorf("address [%s] is invalid", addr)
	}
	return nil
}
