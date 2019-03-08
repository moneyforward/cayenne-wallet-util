package crypto_address_validator

import "testing"

func BenchmarkValidateBTC(b *testing.B) {
	var benchmarks = []struct {
		name        string
		networkType NetworkType
		address     string
	}{
		{"BTC Testnet", NetworkTypeTestNet3, "2MsdiyLjLk9XqHfo5KB9ckptvwm569nX8bm"},
		{"BTC MainNet", NetworkTypeMainNet, "32jExjwLfcZ58EzJFLAPquiqMnZqqKekrL"},
		{"BTC Testnet invalid address", NetworkTypeTestNet3, "asibfcoaibfoia"},
		{"BTC MainNet invalid address", NetworkTypeMainNet, "goaiegbhfoiaebo"},
	}

	b.ResetTimer()
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ValidateBTCAddr(bm.address, bm.networkType)
			}
		})
	}
	b.StopTimer()
}

func BenchmarkValidateBCH(b *testing.B) {
	var benchmarks = []struct {
		name        string
		networkType NetworkType
		address     string
	}{
		{"BCH Testnet", NetworkTypeTestNet3, "bchtest:ppd3wx9aw0m8rgs5ay5q6699app0u745g5th5097p2"},
		{"BCH MainNet", NetworkTypeMainNet, "bitcoincash:pp3w6pvn8ynu7c0waw04m7mv738ysglp6c392ylvql"},
		{"BCH Testnet invalid address", NetworkTypeTestNet3, "asibfcoaibfoia"},
		{"BCH MainNet invalid address", NetworkTypeMainNet, "goaiegbhfoiaebo"},
	}

	b.ResetTimer()
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ValidateBCHAddr(bm.address, bm.networkType)
			}
		})
	}
	b.StopTimer()
}
