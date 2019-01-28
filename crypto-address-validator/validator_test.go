package crypto_address_validator_test

import (
	"testing"

	. "github.com/mf-financial/cayenne-wallet-util/crypto-address-validator"
)

// TestValidateBTCAddressForTestNet (btc/testnet)
func TestValidateBTCAddressForTestNet(t *testing.T) {

	var tests = []struct {
		addr        string
		isErr       bool
		description string
	}{
		{"2NFXSXxw8Fa6P6CSovkdjXE6UF4hupcTHtr", false, "multisig format (for testnet)"},
		{"2NDGkbQTwg2v1zP6yHZw3UJhmsBh9igsSos", false, "multisig format (for testnet)"},
		{"4VHGkbQTGg2vN5P6yHZw3UJhmsBh9igsSos", true, "multisig format with invalid string(for testnet)"},
		{"36oXx9yfVELgpBLYPeyvF7mwjk215MTBPd", true, "multisig format (for mainnet)"},
		{"2MsHnFF6gRvJMv4fSXNLr1Q3zfVp7nYn9GU", false, "BCH Legacy format (for testnet)"}, //FIXME:BCHのアドレスだがこれも通る
		{"1EyWMZofS9wgweqwRnkPVU5YZr5391xM9P", true, "BCH Legacy format (for mainnet)"},
	}

	for _, val := range tests {
		err := ValidateBTCAddr(val.addr, NetworkTypeTestNet3)
		if err != nil && !val.isErr {
			t.Errorf("[Test:%s]\n[address]%s\n Unexpectedly error occorred. %v", val.description, val.addr, err)
		}
		if err == nil && val.isErr {
			t.Errorf("[Test:%s]\n[address]%s\n Error is expected. However nothing happened.", val.description, val.addr)
		}
	}
}

// TestValidateBTCAddressForMainNet (btc/mainnet)
func TestValidateBTCAddressForMainNet(t *testing.T) {

	var tests = []struct {
		addr        string
		isErr       bool
		description string
	}{
		{"15Wsi33cH6drVT9iXx8HNumGG4W6UH3HFN", false, "multisig format (for mainnet)"},
		{"3QKvJCpGKmWrFVwgLgrPZqwS31SJgCBmd6", false, "multisig format (for mainnet)"},
		{"36oXx9yfVELgpBLYPeyvF7mwjk215MTBP4", true, "multisig format with invalid string"},
		{"2NFXSXxw8Fa6P6CSovkdjXE6UF4hupcTHtr", true, "multisig format (for testnet)"},
		{"1EyWMZofS9wgweqwRnkPVU5YZr5391xM9P", false, "BCH Legacy format (for mainnet)"},
	}

	for _, val := range tests {
		err := ValidateBTCAddr(val.addr, NetworkTypeMainNet)
		if err != nil && !val.isErr {
			t.Errorf("[Test:%s]\n[address]%s\n Unexpectedly error occorred. %v", val.description, val.addr, err)
		}
		if err == nil && val.isErr {
			t.Errorf("[Test:%s]\n[address]%s\n Error is expected. However nothing happened.", val.description, val.addr)
		}
	}
}

// TestValidateBCHAddressForTestNet (bch/testnet)
func TestValidateBCHAddressForTestNet(t *testing.T) {

	var tests = []struct {
		addr        string
		isErr       bool
		description string
	}{
		{"bchtest:pqq8hsndcwtgmf0wpvkm4csq2w2z3k6gnqfzza83sp", false, "CashAddr format (for testnet)"},
		{"2MsHnFF6gRvJMv4fSXNLr1Q3zfVp7nYn9GU", false, "Legacy format (for testnet)"},
		{"bchtest:pzupqkxnt2480uev8ms6wt3t3nzssw7pdvuz2lslwp", false, "CashAddr format (for testnet)"},
		{"2NA2TuPV7n2TgsMQKLX6D8FfgC4m8e6NL6n", false, "Legacy format"},
		{"bchtest:pzupqkxnt2480uev8ms6wt3t3nzssw7pdvuz2lslw1", true, "CashAddr format with invalid string (for testnet)"},
		{"2NA2TuPV7n2TgsMQKLX6D8FfgC4m8e6NL61", true, "Legacy format with invalid string (for testnet)"},
		{"bitcoincash:qzv5jmk6g5n0fstd8vzn94zdaq9ptuwjrstwwegql2", true, "CashAddr format (for mainnet)"},
		{"1EyWMZofS9wgweqwRnkPVU5YZr5391xM9P", true, "Legacy format (for mainnet)"},
		{"2NFXSXxw8Fa6P6CSovkdjXE6UF4hupcTHtr", false, "BTC multisig format (for testnet)"}, //FIXME:これはBitcoinのアドレスだが通ってしまうが、OK??
		{"36oXx9yfVELgpBLYPeyvF7mwjk215MTBPd", true, "BTC multisig format (for mainnet)"},
	}

	for _, val := range tests {
		err := ValidateBCHAddr(val.addr, NetworkTypeTestNet3)
		if err != nil && !val.isErr {
			t.Errorf("[Test:%s]\n[address]%s\n Unexpectedly error occorred. %v", val.description, val.addr, err)
		}
		if err == nil && val.isErr {
			t.Errorf("[Test:%s]\n[address]%s\n Error is expected. However nothing happened.", val.description, val.addr)
		}
	}
}

// TestValidateBCHAddressForMainNet (btc/testnet)
func TestValidateBCHAddressForMainNet(t *testing.T) {

	var tests = []struct {
		addr        string
		isErr       bool
		description string
	}{
		{"bitcoincash:qzv5jmk6g5n0fstd8vzn94zdaq9ptuwjrstwwegql2", false, "CashAddr format (for mainnet)"},
		{"1EyWMZofS9wgweqwRnkPVU5YZr5391xM9P", false, "Legacy format (for mainnet)"},
		{"bitcoincash:qzv5jmk6g5n0fstd8vzn94zdaq9ptuwjrstwwegql3", true, "CashAddr format with invalid string (for mainnet)"},
		{"1EyWMZofS9wgweqwRnkPVU5YZr5391xM93", true, "Legacy format with invalid string (for mainnet)"},
		{"bchtest:pqq8hsndcwtgmf0wpvkm4csq2w2z3k6gnqfzza83sp", true, "CashAddr format (for testnet)"},
		{"2MsHnFF6gRvJMv4fSXNLr1Q3zfVp7nYn9GU", true, "Legacy format (for testnet)"},
		{"2NFXSXxw8Fa6P6CSovkdjXE6UF4hupcTHtr", true, "BTC multisig format (for testnet)"},
		{"36oXx9yfVELgpBLYPeyvF7mwjk215MTBPd", false, "BTC multisig format (for mainnet)"}, //FIXME:これはBitcoinのアドレスだが通ってしまうが、OK??
	}

	for _, val := range tests {
		err := ValidateBCHAddr(val.addr, NetworkTypeMainNet)
		if err != nil && !val.isErr {
			t.Errorf("[Test:%s]\n[address]%s\n Unexpectedly error occorred. %v", val.description, val.addr, err)
		}
		if err == nil && val.isErr {
			t.Errorf("[Test:%s]\n[address]%s\n Error is expected. However nothing happened.", val.description, val.addr)
		}
	}
}
